package tempconv

// CToF konwertuje temperaturę w stopniach Celsjusza na stopnie Fahrenheita.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC konwertuje temperaturę w stopniach Fahrenheita na stopnie Celsjusza.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// MToF konwertuje metry na stopy.
func MToF(m Meter) Foot { return Foot(m * 3.28) }

// FToM konwertuje stopy na metry.
func FToM(f Foot) Meter { return Meter(f / 3.28) }

// PToK konwertuje funty na kilogramy.
func PToK(p Pound) Kg { return Kg(p / 2.2) }

// KToP konwertuje kilogramy na funty.
func KToP(k Kg) Pound { return Pound(k * 2.2) }
