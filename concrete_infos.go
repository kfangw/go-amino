package amino

func (cdc *Codec) ConcreteInfos() []*TypeInfo {
	typeInfos := make([]*TypeInfo, 0, len(cdc.concreteInfos))
	for _, ti := range cdc.concreteInfos {
		typeInfos = append(typeInfos, ti)
	}
	return typeInfos
}
