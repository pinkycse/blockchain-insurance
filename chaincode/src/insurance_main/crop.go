package main

//==============================================================================================================================
//	Crop - Defines the structure for a Crop object.
//==============================================================================================================================
type Crop struct {
	Id		string			`json:"id"`
	Type		string			`json:"type"`
	Details		CropDetails		`json:"details"`
}

//==============================================================================================================================
//	CropDetails - Defines the structure for a CropDetails object.
//==============================================================================================================================
type CropDetails struct {
	Name			string		`json:"name"`
	Type			string		`json:"type"`
	Duration		int		`json:"duration"`
}

//=================================================================================================================================
//	 New Crop	-	Constructs a new crop
//=================================================================================================================================
func NewCrop(name string, type string, duration int ) (Crop) {
	var crop Crop

	crop.Id = name
	crop.Type = type

	crop.Details.Name = name
	crop.Details.Type = type
	crop.Details.Duration = duration

	return crop
}
