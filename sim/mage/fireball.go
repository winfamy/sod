package mage

import (
	"fmt"
	"time"

	"github.com/wowsims/sod/sim/core"
)

const FireballRanks = 12

var FireballSpellId = [FireballRanks + 1]int32{0, 133, 143, 145, 3140, 8400, 8401, 8402, 10148, 10149, 10150, 10151, 25306}
var FireballBaseDamage = [FireballRanks + 1][]float64{{0}, {16, 25}, {34, 49}, {57, 77}, {89, 122}, {140, 189}, {207, 274}, {264, 345}, {328, 425}, {398, 512}, {488, 623}, {561, 715}, {596, 760}}
var FireballDotDamage = [FireballRanks + 1]float64{0, 2, 3, 6, 12, 20, 28, 32, 40, 52, 60, 72, 76}
var FireballSpellCoeff = [FireballRanks + 1]float64{0, .123, .271, .5, .793, 1, 1, 1, 1, 1, 1, 1, 1}
var FireballCastTime = [FireballRanks + 1]int32{0, 1500, 2000, 2500, 3000, 3500, 3500, 3500, 3500, 3500, 3500, 3500, 3500}
var FireballManaCost = [FireballRanks + 1]float64{0, 30, 45, 65, 95, 140, 185, 220, 260, 305, 350, 395, 410}
var FireballLevel = [FireballRanks + 1]int{0, 1, 6, 12, 18, 24, 30, 36, 42, 48, 54, 60, 60}

func (mage *Mage) registerFireballSpell() {
	mage.Fireball = make([]*core.Spell, FireballRanks+1)

	for rank := 1; rank <= FireballRanks; rank++ {
		config := mage.getFireballBaseConfig(rank)

		if config.RequiredLevel <= int(mage.Level) {
			mage.Fireball[rank] = mage.GetOrRegisterSpell(config)
		}
	}
}

func (mage *Mage) getFireballBaseConfig(rank int) core.SpellConfig {
	spellId := FireballSpellId[rank]
	baseDamageLow := FireballBaseDamage[rank][0]
	baseDamageHigh := FireballBaseDamage[rank][1]
	baseDotDamage := FireballDotDamage[rank]
	spellCoeff := FireballSpellCoeff[rank]
	castTime := FireballCastTime[rank]
	manaCost := FireballManaCost[rank]
	level := FireballLevel[rank]

	ticks := int32(4)

	return core.SpellConfig{
		ActionID:      core.ActionID{SpellID: spellId},
		SpellSchool:   core.SpellSchoolFire,
		ProcMask:      core.ProcMaskSpellDamage,
		Flags:         core.SpellFlagAPL,
		RequiredLevel: level,
		Rank:          rank,
		MissileSpeed:  24,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:      core.GCDDefault,
				CastTime: time.Millisecond*time.Duration(castTime) - time.Millisecond*100*time.Duration(mage.Talents.ImprovedFireball),
			},
		},
		Dot: core.DotConfig{
			Aura: core.Aura{
				Label: fmt.Sprintf("Fireball (Rank %d)", rank),
			},
			NumberOfTicks: ticks,
			TickLength:    time.Second * 2,
			OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, _ bool) {
				dot.SnapshotBaseDamage = baseDotDamage / float64(ticks)
				dot.SnapshotAttackerMultiplier = dot.Spell.AttackerDamageMultiplier(dot.Spell.Unit.AttackTables[target.UnitIndex][dot.Spell.CastType])
			},
			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
			},
		},

		DamageMultiplier: 1,
		CritMultiplier:   mage.DefaultSpellCritMultiplier(),
		ThreatMultiplier: 1 - (0.15 * float64(mage.Talents.BurningSoul)),

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := sim.Roll(baseDamageLow, baseDamageHigh) + spellCoeff*spell.SpellDamage()
			result := spell.CalcDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)
			spell.WaitTravelTime(sim, func(sim *core.Simulation) {
				if result.Landed() {
					spell.DealDamage(sim, result)
					spell.Dot(target).Apply(sim)
				}
			})
		},
	}
}
