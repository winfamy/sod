import { CURRENT_PHASE, LEVEL_THRESHOLDS, Phase } from './other';

export const MAX_CHARACTER_LEVEL = LEVEL_THRESHOLDS[Phase.Phase5];
export const MAX_TALENT_POINTS = MAX_CHARACTER_LEVEL - 9;
export const CURRENT_LEVEL_CAP = LEVEL_THRESHOLDS[CURRENT_PHASE];
export const BOSS_LEVEL = MAX_CHARACTER_LEVEL + 3;

export const EXPERTISE_PER_QUARTER_PERCENT_REDUCTION = 32.79 / 4;
export const MELEE_CRIT_RATING_PER_CRIT_CHANCE = 1;
export const MELEE_HIT_RATING_PER_HIT_CHANCE = 1;
export const ARMOR_PEN_PER_PERCENT_ARMOR = 13.99;

export const SPELL_CRIT_RATING_PER_CRIT_CHANCE = 1;
export const SPELL_HIT_RATING_PER_HIT_CHANCE = 1;

export const HASTE_RATING_PER_HASTE_PERCENT = 1;

export const DEFENSE_RATING_PER_DEFENSE = 4.92;
export const MISS_DODGE_PARRY_BLOCK_CRIT_CHANCE_PER_DEFENSE = 0.04;
export const BLOCK_RATING_PER_BLOCK_CHANCE = 1;
export const DODGE_RATING_PER_DODGE_CHANCE = 1;
export const PARRY_RATING_PER_PARRY_CHANCE = 1;
export const RESILIENCE_RATING_PER_CRIT_REDUCTION_CHANCE = 94.27;
export const RESILIENCE_RATING_PER_CRIT_DAMAGE_REDUCTION_PERCENT = 94.27 / 2.2;
