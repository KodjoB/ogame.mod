package ogame

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEnergyTechnologyConstructionTime(t *testing.T) {
	mm := newEnergyTechnology()
	ct := mm.ConstructionTime(5, 7, Facilities{ResearchLab: 3})
	assert.Equal(t, 1645*time.Second, ct)
}
