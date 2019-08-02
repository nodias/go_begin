//go:generate stringer -type=Pill
package painkiller

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofin
	Paracetamol
	Acetaminophen = Paracetamol
)

//func (p Pill) String() string{
//	switch  p {
//	case Placebo:
//		return "Placebo"
//	case Aspirin:
//		return "Aspirin:"
//	case Ibuprofin:
//		return "Ibuprofin"
//	case Paracetamol: // == Acetaminophen
//		return "Paracetamol"
//	}
//	return fmt.Sprintf("Pill(%d)", p)
//}