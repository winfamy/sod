package mage

import (
	"time"

	"github.com/wowsims/sod/sim/core"
	"github.com/wowsims/sod/sim/core/proto"
	"github.com/wowsims/sod/sim/core/stats"
)

// TODO: Classic verify Arcane Blast rune numbers
// https://www.wowhead.com/classic/news/patch-1-15-build-52124-ptr-datamining-season-of-discovery-runes-336044#news-post-336044
// https://www.wowhead.com/classic/spell=400574/arcane-blast
func (mage *Mage) registerArcaneBlastSpell() {
	if !mage.HasRune(proto.MageRune_RuneHandsArcaneBlast) {
		return
	}

	level := float64(mage.GetCharacter().Level)
	baseCalc := (13.828124 + 0.018012*level + 0.044141*level*level)
	baseLowDamage := baseCalc * 4.53
	baseHighDamage := baseCalc * 5.27
	spellCoeff := .714
	castTime := time.Millisecond * 2500
	manaCost := .07

	mage.ArcaneBlastAura = mage.GetOrRegisterAura(core.Aura{
		Label:     "Arcane Blast Aura",
		ActionID:  core.ActionID{SpellID: 400573},
		Duration:  time.Second * 6,
		MaxStacks: 4,
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			aura.Refresh(sim)
			mage.ArcaneBlast.CostMultiplier = 1.75 * float64(newStacks)
			mage.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexArcane] /= 1 + .15*float64(oldStacks)
			mage.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexArcane] *= 1 + .15*float64(newStacks)
		},
		OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
			if !spell.SpellSchool.Matches(core.SpellSchoolArcane) || !spell.Flags.Matches(SpellFlagMage) || spell == mage.ArcaneBlast {
				return
			}

			aura.Deactivate(sim)
		},
	})

	mage.ArcaneBlast = mage.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 400574},
		SpellCode:   SpellCode_MageArcaneBlast,
		SpellSchool: core.SpellSchoolArcane,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       SpellFlagMage | core.SpellFlagAPL,

		ManaCost: core.ManaCostOptions{
			BaseCost: manaCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:      core.GCDDefault,
				CastTime: castTime,
			},
		},

		DamageMultiplier: 1,
		CritMultiplier:   mage.MageCritMultiplier(0),
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := sim.Roll(baseLowDamage, baseHighDamage) + spellCoeff*spell.SpellDamage()
			result := spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)

			if result.Landed() {
				if !mage.ArcaneBlastAura.IsActive() {
					mage.ArcaneBlastAura.Activate(sim)
				}
				if mage.ArcaneBlastAura.GetStacks() == mage.ArcaneBlastAura.MaxStacks {
					mage.ArcaneBlastAura.Refresh(sim)
				}
				mage.ArcaneBlastAura.AddStack(sim)
			}
		},
	})
}
