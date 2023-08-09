package builder

type iBuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() house
}

func GetBuiler(builerType string) iBuilder {
	if builerType == "normal" {
		return &normalBuilder{}
	}
	if builerType == "igloo" {
		return &iglooBuilder{}
	}
	return nil
}
