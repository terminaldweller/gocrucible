package main

type OSMKeyValuePair struct {
	Key   string
	Value string
}

// Reference: https://wiki.openstreetmap.org/wiki/Map_Features
func makePriorityMap() []OSMKeyValuePair {
	result := []OSMKeyValuePair{
		{"place", "country"},
		{"place", "province"},
		{"place", "state"},
		{"place", "county"},
		{"place", "city"},
		{"place", "town"},
		{"place", "municipality"},
		{"place", "district"},
		{"place", "region"},
		{"place", "village"},
		{"place", "borough"},
		{"place", "suburb"},
		{"place", "quarter"},
		{"place", "neighbourhood"},
		{"place", "city_block"},
		{"place", "square"},
		{"route", "road"},
		{"building", "hospital"},
		{"aeroway", "aerodrome"},
		{"place", ""},
		{"aeroway", "helipad"},
		{"aeroway", "heliport"},
		{"aeroway", "taxiway"},
		{"aeroway", "terminal"},
		{"aeroway", ""},
		{"amenity", "baby_hatch"},
		{"amenity", "clinic"},
		{"amenity", "dentist"},
		{"amenity", "doctors"},
		{"amenity", "hospital"},
		{"amenity", "nursing_home"},
		{"amenity", "pharmacy"},
		{"amenity", "social_facility"},
		{"amenity", "veteniray"},
		{"amenity", "community_centre"},
		{"amenity", "sanitary_dump_station"},
		{"amenity", ""},
		{"barrier", ""},
		{"boundary", ""},
		{"building", ""},
		{"craft", "optician"},
		{"emergency", ""},
		{"shop", "medical_supply"},
		{"shop", "optician"},
		{"highway", ""},
		{"highway", "road"},
		{"highway", "path"},
		{"footway", "sidewalk"},
		{"footway", "crossing"},
		{"sidewalk", "both"},
		{"sidewalk", "left"},
		{"sidewalk", "right"},
		{"sidewalk", "no"},
		{"highway", "cycleway"},
		{"cycleway", "lane"},
		{"cycleway", ""},
		{"busway", "lane"},
		{"highway", "proposed"},
		{"highway", "construction"},
		{"highway", "bus_stop"},
		{"highway", "crossing"},
		{"highway", "elevator"},
		{"highway", "emergency_bay"},
		{"emergency", "phone"},
		{"leisure", ""},
		{"man_made", ""},
		{"military", ""},
		{"natural", ""},
		{"office", ""},
		{"power", ""},
		{"line", ""},
		{"public_transport", ""},
		{"railway", ""},
		{"bridge", ""},
		{"service", ""},
		{"route", ""},
		{"shop", ""},
		{"sport", ""},
		{"telecom", ""},
		{"aerialway", ""},
		{"tourism", ""},
		{"waterway", ""},
		// {"building", "hotel"},
		// {"building", "apartments"},
		// {"building", "dormitory"},
		// {"building", "house"},
		// {"building", "residential"},
		// {"building", "commercial"},
		// {"building", "industrial"},
		// {"building", "office"},
		// {"building", "supermarket"},
		// {"building", "warehouse"},
		// {"building", "cathedral"},
		// {"building", "chapel"},
		// {"building", "church"},
		// {"building", "mosque"},
		// {"building", "religious"},
		// {"building", "shrine"},
		// {"building", "synagogue"},
		// {"building", "temple"},
		// {"building", "fire_station"},
		// {"building", "government"},
		// {"building", "train_station"},
		// {"building", "transportation"},
		// {"building", "kindergarten"},
		// {"building", "school"},
		// {"building", "university"},
		// {"building", "college"},
		// {"building", "greenhouse"},
		// {"building", "parking"},
		// {"building", "bridge"},
		// craft keys are out of scope
		// {"emergency", "ambulance_station"},
		// {"emergency", "defibrillator"},
		// {"emergency", "landing_site"},
		// {"emergency", "emergency_ward_entrace"},
		// {"emergency", "dry_riser_inlet"},
		// {"emergency", "fire_alram_box"},
		// {"emergency", "fire_extinguisher"},
		// {"emergency", "fire_hose"},
		// {"emergency", "fire_hydrant"},
		// {"emergency", "water_tank"},
		// {"emergency", "suction_point"},
		// {"emergency", "lifeguard"},
		// {"emergency", "lifeguard_base"},
		// {"emergency", "lifeguard_tower"},
		// {"emergency", "lifeguard_platform"},
		// {"emergency", "life_ring"},
		// {"emergency", "assembly_point"},
		// {"emergency", "phone"},
		// {"emergency", "siren"},
		// {"addr:housenumber", ""},
		// {"addr:housename", ""},
		// {"addr:flats", ""},
		// {"addr:conscriptionnumber", ""},
		// {"addr:street", ""},
		// {"addr:place", ""},
		// {"addr:postcode", ""},
		// {"addr:city", ""},
		// {"addr:country", ""},
		// {"addr:full", ""},
		// {"addr:hamlet", ""},
		// {"addr:suburb", ""},
		// {"addr:subdistrict", ""},
		// {"addr:district", ""},
		// {"addr:province", ""},
		// {"addr:state", ""},
		// {"addr:interpolation", ""},
		// {"addr:inclusion", ""},
	}

	return result
}