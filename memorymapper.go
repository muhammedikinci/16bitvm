package main

import (
	"errors"
)

type Regions []Region

type MemoryMapper struct {
	Regions Regions
}

type Region struct {
	Device IMemory
	Start  uint16
	End    uint16
	Remap  bool
}

func NewMemoryMapper() *MemoryMapper {
	return &MemoryMapper{Regions: Regions{}}
}

func (m *MemoryMapper) Map(region Region) Regions {
	prevRegions := m.Regions
	m.Regions = m.Regions.unshift(region)
	return prevRegions
}

func (m *MemoryMapper) FindRegion(address uint16) (Region, error) {
	for i := 0; i < len(m.Regions); i++ {
		if m.Regions[i].Start <= address && m.Regions[i].End >= address {
			return m.Regions[i], nil
		}
	}
	return Region{}, errors.New("there is not region for specified address")
}

func (m *MemoryMapper) Get16(address uint16) uint16 {
	region, err := m.FindRegion(address)
	if err != nil {
		return 0x00
	}

	retAddr := address

	if region.Remap {
		retAddr = address - region.Start
	}

	return region.Device.Get16(retAddr)
}

func (m *MemoryMapper) Get8(address uint16) uint8 {
	region, err := m.FindRegion(address)
	if err != nil {
		return 0x00
	}

	retAddr := address

	if region.Remap {
		retAddr = address - region.Start
	}

	return region.Device.Get8(retAddr)
}

func (m *MemoryMapper) Set16(address, value uint16) {
	region, err := m.FindRegion(address)
	if err != nil {
		return
	}

	retAddr := address

	if region.Remap {
		retAddr = address - region.Start
	}

	region.Device.Set16(retAddr, value)
}

func (m *MemoryMapper) Set8(address uint16, value uint8) {
	region, err := m.FindRegion(address)
	if err != nil {
		return
	}

	retAddr := address

	if region.Remap {
		retAddr = address - region.Start
	}

	region.Device.Set8(retAddr, value)
}

func (r Regions) unshift(val Region) Regions {
	return append(Regions{val}, r...)
}
