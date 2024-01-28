import { Spec } from '../core/proto/common.js';
import { ActionId } from '../core/proto_utils/action_id.js';

import {
	ProtectionWarrior_Rotation_DemoShoutChoice as DemoShoutChoice,
	ProtectionWarrior_Rotation_SpellOption as SpellOption,
	ProtectionWarrior_Rotation_ThunderClapChoice as ThunderClapChoice,
	WarriorShout
} from '../core/proto/warrior.js';

import * as InputHelpers from '../core/components/input_helpers.js';

// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.

export const StartingRage = InputHelpers.makeSpecOptionsNumberInput<Spec.SpecProtectionWarrior>({
	fieldName: 'startingRage',
	label: 'Starting Rage',
	labelTooltip: 'Initial rage at the start of each iteration.',
});

export const ShoutPicker = InputHelpers.makeSpecOptionsEnumIconInput<Spec.SpecProtectionWarrior, WarriorShout>({
	fieldName: 'shout',
	values: [
		{ color: 'c79c6e', value: WarriorShout.WarriorShoutNone },
		{ actionId: ActionId.fromSpellId(47436), value: WarriorShout.WarriorShoutBattle },
		{ actionId: ActionId.fromSpellId(469), value: WarriorShout.WarriorShoutCommanding },
	],
});

export const ShatteringThrow = InputHelpers.makeSpecOptionsBooleanIconInput<Spec.SpecProtectionWarrior>({
	fieldName: 'useShatteringThrow',
	actionId: ActionId.fromSpellId(64382),
});

export const ProtectionWarriorRotationConfig = {
	inputs: [
		InputHelpers.makeCustomRotationInput<Spec.SpecProtectionWarrior, SpellOption>({
			fieldName: 'customRotation',
			numColumns: 3,
			values: [
				{ actionId: ActionId.fromSpellId(57823), value: SpellOption.Revenge },
				{ actionId: ActionId.fromSpellId(47488), value: SpellOption.ShieldSlam },
				{ actionId: ActionId.fromSpellId(47440), value: SpellOption.Shout },
				{ actionId: ActionId.fromSpellId(47502), value: SpellOption.ThunderClap },
				{ actionId: ActionId.fromSpellId(25203), value: SpellOption.DemoralizingShout },
				{ actionId: ActionId.fromSpellId(47486), value: SpellOption.MortalStrike },
				{ actionId: ActionId.fromSpellId(47498), value: SpellOption.Devastate },
				{ actionId: ActionId.fromSpellId(47467), value: SpellOption.SunderArmor },
				{ actionId: ActionId.fromSpellId(12809), value: SpellOption.ConcussionBlow },
				{ actionId: ActionId.fromSpellId(46968), value: SpellOption.Shockwave },
			],
		}),
		InputHelpers.makeRotationNumberInput<Spec.SpecProtectionWarrior>({
			fieldName: 'hsRageThreshold',
			label: 'HS rage threshold',
			labelTooltip: 'Heroic Strike when rage is above:',
		}),
		InputHelpers.makeRotationBooleanInput<Spec.SpecProtectionWarrior>({
			fieldName: 'prioSslamOnShieldBlock',
			label: 'Prio SSlam on Shield Block',
			labelTooltip: 'The rotation code will prio SSlam over Revenge during active shield block windows.',
		}),
		InputHelpers.makeRotationEnumInput<Spec.SpecProtectionWarrior>({
			fieldName: 'demoShoutChoice',
			label: 'Demo Shout',
			values: [
				{ name: 'Never', value: DemoShoutChoice.DemoShoutChoiceNone },
				{ name: 'Maintain Debuff', value: DemoShoutChoice.DemoShoutChoiceMaintain },
				{ name: 'Filler', value: DemoShoutChoice.DemoShoutChoiceFiller },
			],
		}),
		InputHelpers.makeRotationEnumInput<Spec.SpecProtectionWarrior>({
			fieldName: 'thunderClapChoice',
			label: 'Thunder Clap',
			values: [
				{ name: 'Never', value: ThunderClapChoice.ThunderClapChoiceNone },
				{ name: 'Maintain Debuff', value: ThunderClapChoice.ThunderClapChoiceMaintain },
				{ name: 'On CD', value: ThunderClapChoice.ThunderClapChoiceOnCD },
			],
		}),
	],
};
