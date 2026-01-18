// Package pokeapi provides types and utilities for interacting with the PokeAPI.
package pokeapi

// LocationAreasResponse represents the paginated response from the location-area list endpoint.
type LocationAreasResponse struct {
	Count    int             `json:"count"`
	Next     *string         `json:"next"`
	Previous *string         `json:"previous"`
	Results  []NamedResource `json:"results"`
}

// LocationAreaResponse represents the response from a specific location-area endpoint.
// It contains detailed information about Pokemon encounters in that area.
type LocationAreaResponse struct {
	ID                   int                   `json:"id"`
	Name                 string                `json:"name"`
	GameIndex            int                   `json:"game_index"`
	Location             NamedResource         `json:"location"`
	Names                []LocalizedName       `json:"names"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

// NamedResource is a common structure for API resources with a name and URL.
type NamedResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// LocalizedName represents a name in a specific language.
type LocalizedName struct {
	Name     string        `json:"name"`
	Language NamedResource `json:"language"`
}

// EncounterMethodRate describes encounter method rates for the location area.
type EncounterMethodRate struct {
	EncounterMethod NamedResource   `json:"encounter_method"`
	VersionDetails  []VersionDetail `json:"version_details"`
}

// VersionDetail contains version-specific rate information.
type VersionDetail struct {
	Rate    int           `json:"rate"`
	Version NamedResource `json:"version"`
}

// PokemonEncounter describes a Pokemon that can be encountered in the location.
type PokemonEncounter struct {
	Pokemon        NamedResource           `json:"pokemon"`
	VersionDetails []VersionEncounterGroup `json:"version_details"`
}

// VersionEncounterGroup contains encounter details for a specific game version.
type VersionEncounterGroup struct {
	Version          NamedResource     `json:"version"`
	MaxChance        int               `json:"max_chance"`
	EncounterDetails []EncounterDetail `json:"encounter_details"`
}

// EncounterDetail describes the specifics of how a Pokemon encounter occurs.
type EncounterDetail struct {
	MinLevel        int           `json:"min_level"`
	MaxLevel        int           `json:"max_level"`
	Chance          int           `json:"chance"`
	Method          NamedResource `json:"method"`
	ConditionValues []any         `json:"condition_values"`
}

// Pokemon represents detailed information about a specific Pokemon from the PokeAPI.
type Pokemon struct {
	ID                     int              `json:"id"`
	Name                   string           `json:"name"`
	BaseExperience         int              `json:"base_experience"`
	Height                 int              `json:"height"`
	Weight                 int              `json:"weight"`
	IsDefault              bool             `json:"is_default"`
	Order                  int              `json:"order"`
	Species                NamedResource    `json:"species"`
	LocationAreaEncounters string           `json:"location_area_encounters"`
	Abilities              []PokemonAbility `json:"abilities"`
	Forms                  []NamedResource  `json:"forms"`
	GameIndices            []GameIndex      `json:"game_indices"`
	HeldItems              []HeldItem       `json:"held_items"`
	Moves                  []PokemonMove    `json:"moves"`
	Stats                  []PokemonStat    `json:"stats"`
	Types                  []PokemonType    `json:"types"`
	PastTypes              []PastType       `json:"past_types"`
	PastAbilities          []PastAbility    `json:"past_abilities"`
	Sprites                PokemonSprites   `json:"sprites"`
	Cries                  PokemonCries     `json:"cries"`
}

// PokemonAbility represents an ability a Pokemon can have.
type PokemonAbility struct {
	IsHidden bool          `json:"is_hidden"`
	Slot     int           `json:"slot"`
	Ability  NamedResource `json:"ability"`
}

// GameIndex represents a game index for a Pokemon in a specific version.
type GameIndex struct {
	GameIndex int           `json:"game_index"`
	Version   NamedResource `json:"version"`
}

// HeldItem represents an item a Pokemon may hold in the wild.
type HeldItem struct {
	Item           NamedResource     `json:"item"`
	VersionDetails []HeldItemVersion `json:"version_details"`
}

// HeldItemVersion contains version-specific held item rarity.
type HeldItemVersion struct {
	Rarity  int           `json:"rarity"`
	Version NamedResource `json:"version"`
}

// PokemonMove represents a move a Pokemon can learn.
type PokemonMove struct {
	Move                NamedResource       `json:"move"`
	VersionGroupDetails []MoveVersionDetail `json:"version_group_details"`
}

// MoveVersionDetail contains details about how a move is learned in a version group.
type MoveVersionDetail struct {
	LevelLearnedAt  int           `json:"level_learned_at"`
	Order           int           `json:"order"`
	VersionGroup    NamedResource `json:"version_group"`
	MoveLearnMethod NamedResource `json:"move_learn_method"`
}

// PokemonStat represents a Pokemon's base stat.
type PokemonStat struct {
	BaseStat int           `json:"base_stat"`
	Effort   int           `json:"effort"`
	Stat     NamedResource `json:"stat"`
}

// PokemonType represents a Pokemon's type and slot.
type PokemonType struct {
	Slot int           `json:"slot"`
	Type NamedResource `json:"type"`
}

// PastType represents a Pokemon's type in previous generations.
type PastType struct {
	Generation NamedResource `json:"generation"`
	Types      []PokemonType `json:"types"`
}

// PastAbility represents a Pokemon's abilities in previous generations.
type PastAbility struct {
	Generation NamedResource      `json:"generation"`
	Abilities  []PastAbilityEntry `json:"abilities"`
}

// PastAbilityEntry represents a single past ability entry.
type PastAbilityEntry struct {
	Ability  any  `json:"ability"`
	IsHidden bool `json:"is_hidden"`
	Slot     int  `json:"slot"`
}

// PokemonCries contains URLs to Pokemon cry audio files.
type PokemonCries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

// PokemonSprites contains all sprite image URLs for a Pokemon.
type PokemonSprites struct {
	BackDefault      string         `json:"back_default"`
	BackFemale       any            `json:"back_female"`
	BackShiny        string         `json:"back_shiny"`
	BackShinyFemale  any            `json:"back_shiny_female"`
	FrontDefault     string         `json:"front_default"`
	FrontFemale      any            `json:"front_female"`
	FrontShiny       string         `json:"front_shiny"`
	FrontShinyFemale any            `json:"front_shiny_female"`
	Other            OtherSprites   `json:"other"`
	Versions         VersionSprites `json:"versions"`
}

// OtherSprites contains alternative sprite sources.
type OtherSprites struct {
	DreamWorld      DreamWorldSprites      `json:"dream_world"`
	Home            HomeSprites            `json:"home"`
	OfficialArtwork OfficialArtworkSprites `json:"official-artwork"`
	Showdown        ShowdownSprites        `json:"showdown"`
}

// DreamWorldSprites contains Dream World sprite URLs.
type DreamWorldSprites struct {
	FrontDefault string `json:"front_default"`
	FrontFemale  any    `json:"front_female"`
}

// HomeSprites contains Pokemon Home sprite URLs.
type HomeSprites struct {
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}

// OfficialArtworkSprites contains official artwork URLs.
type OfficialArtworkSprites struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

// ShowdownSprites contains Pokemon Showdown sprite URLs.
type ShowdownSprites struct {
	BackDefault      string `json:"back_default"`
	BackFemale       any    `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  any    `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}

// VersionSprites contains sprites organized by game generation.
type VersionSprites struct {
	GenerationI    GenerationISprites    `json:"generation-i"`
	GenerationII   GenerationIISprites   `json:"generation-ii"`
	GenerationIII  GenerationIIISprites  `json:"generation-iii"`
	GenerationIV   GenerationIVSprites   `json:"generation-iv"`
	GenerationV    GenerationVSprites    `json:"generation-v"`
	GenerationVI   GenerationVISprites   `json:"generation-vi"`
	GenerationVII  GenerationVIISprites  `json:"generation-vii"`
	GenerationVIII GenerationVIIISprites `json:"generation-viii"`
}

// GenerationISprites contains Generation I game sprites.
type GenerationISprites struct {
	RedBlue RedBlueSprites `json:"red-blue"`
	Yellow  YellowSprites  `json:"yellow"`
}

// RedBlueSprites contains Red/Blue version sprites.
type RedBlueSprites struct {
	BackDefault  string `json:"back_default"`
	BackGray     string `json:"back_gray"`
	FrontDefault string `json:"front_default"`
	FrontGray    string `json:"front_gray"`
}

// YellowSprites contains Yellow version sprites.
type YellowSprites struct {
	BackDefault  string `json:"back_default"`
	BackGray     string `json:"back_gray"`
	FrontDefault string `json:"front_default"`
	FrontGray    string `json:"front_gray"`
}

// GenerationIISprites contains Generation II game sprites.
type GenerationIISprites struct {
	Crystal CrystalSprites `json:"crystal"`
	Gold    GoldSprites    `json:"gold"`
	Silver  SilverSprites  `json:"silver"`
}

// CrystalSprites contains Crystal version sprites.
type CrystalSprites struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

// GoldSprites contains Gold version sprites.
type GoldSprites struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

// SilverSprites contains Silver version sprites.
type SilverSprites struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

// GenerationIIISprites contains Generation III game sprites.
type GenerationIIISprites struct {
	Emerald          EmeraldSprites          `json:"emerald"`
	FireredLeafgreen FireredLeafgreenSprites `json:"firered-leafgreen"`
	RubySapphire     RubySapphireSprites     `json:"ruby-sapphire"`
}

// EmeraldSprites contains Emerald version sprites.
type EmeraldSprites struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

// FireredLeafgreenSprites contains FireRed/LeafGreen version sprites.
type FireredLeafgreenSprites struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

// RubySapphireSprites contains Ruby/Sapphire version sprites.
type RubySapphireSprites struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

// GenerationIVSprites contains Generation IV game sprites.
type GenerationIVSprites struct {
	DiamondPearl        DiamondPearlSprites        `json:"diamond-pearl"`
	HeartgoldSoulsilver HeartgoldSoulsilverSprites `json:"heartgold-soulsilver"`
	Platinum            PlatinumSprites            `json:"platinum"`
}

// DiamondPearlSprites contains Diamond/Pearl version sprites.
type DiamondPearlSprites struct {
	BackDefault      string `json:"back_default"`
	BackFemale       any    `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  any    `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}

// HeartgoldSoulsilverSprites contains HeartGold/SoulSilver version sprites.
type HeartgoldSoulsilverSprites struct {
	BackDefault      string `json:"back_default"`
	BackFemale       any    `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  any    `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}

// PlatinumSprites contains Platinum version sprites.
type PlatinumSprites struct {
	BackDefault      string `json:"back_default"`
	BackFemale       any    `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  any    `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}

// GenerationVSprites contains Generation V game sprites.
type GenerationVSprites struct {
	BlackWhite BlackWhiteSprites `json:"black-white"`
}

// BlackWhiteSprites contains Black/White version sprites.
type BlackWhiteSprites struct {
	Animated         AnimatedSprites `json:"animated"`
	BackDefault      string          `json:"back_default"`
	BackFemale       any             `json:"back_female"`
	BackShiny        string          `json:"back_shiny"`
	BackShinyFemale  any             `json:"back_shiny_female"`
	FrontDefault     string          `json:"front_default"`
	FrontFemale      any             `json:"front_female"`
	FrontShiny       string          `json:"front_shiny"`
	FrontShinyFemale any             `json:"front_shiny_female"`
}

// AnimatedSprites contains animated sprite URLs.
type AnimatedSprites struct {
	BackDefault      string `json:"back_default"`
	BackFemale       any    `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  any    `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}

// GenerationVISprites contains Generation VI game sprites.
type GenerationVISprites struct {
	OmegarubyAlphasapphire OmegarubyAlphasapphireSprites `json:"omegaruby-alphasapphire"`
	XY                     XYSprites                     `json:"x-y"`
}

// OmegarubyAlphasapphireSprites contains Omega Ruby/Alpha Sapphire version sprites.
type OmegarubyAlphasapphireSprites struct {
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}

// XYSprites contains X/Y version sprites.
type XYSprites struct {
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}

// GenerationVIISprites contains Generation VII game sprites.
type GenerationVIISprites struct {
	Icons             IconSprites              `json:"icons"`
	UltraSunUltraMoon UltraSunUltraMoonSprites `json:"ultra-sun-ultra-moon"`
}

// IconSprites contains small icon sprites.
type IconSprites struct {
	FrontDefault string `json:"front_default"`
	FrontFemale  any    `json:"front_female"`
}

// UltraSunUltraMoonSprites contains Ultra Sun/Ultra Moon version sprites.
type UltraSunUltraMoonSprites struct {
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}

// GenerationVIIISprites contains Generation VIII game sprites.
type GenerationVIIISprites struct {
	Icons IconSprites `json:"icons"`
}
