package input

// Translator is a minimal placeholder for the Kotlin Translator. For the
// Go prototype we provide a simple map-based translator.
type Translator struct {
	translationDir string
	mapData        map[string]string
}

func NewTranslator(translationDir string) *Translator {
	return &Translator{translationDir: translationDir, mapData: make(map[string]string)}
}

func (t *Translator) AddKey(key, value string) { t.mapData[key] = value }
func (t *Translator) Translate(key string) string {
	if v, ok := t.mapData[key]; ok {
		return v
	}
	return key
}
