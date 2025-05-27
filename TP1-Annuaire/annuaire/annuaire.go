package annuaire

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

const fichier = "contacts.json"

type Contact struct {
	Nom string `json:"nom"`
	Tel string `json:"tel"`
}

type Annuaire struct {
	Contacts map[string]Contact
}

func NewAnnuaire() *Annuaire {
	return &Annuaire{
		Contacts: make(map[string]Contact),
	}
}

func (a *Annuaire) Charger() error {
	data, err := os.ReadFile(fichier)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &a.Contacts)
}

func (a *Annuaire) Sauvegarder() error {
	data, err := json.MarshalIndent(a.Contacts, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fichier, data, 0644)
}

func (a *Annuaire) getNextID() string {
	max := 0
	for id := range a.Contacts {
		if n, err := strconv.Atoi(id); err == nil && n > max {
			max = n
		}
	}
	return strconv.Itoa(max + 1)
}

func (a *Annuaire) Rechercher(nom string) bool {
	for _, c := range a.Contacts {
		if c.Nom == nom {
			return true
		}
	}
	return false
}

func isValidPhoneNumber(tel string) bool {
	re := regexp.MustCompile(`^\d{10}$`)
	return re.MatchString(tel)
}

func (a *Annuaire) Ajouter(nom, tel string) error {
	if a.Rechercher(nom) {
		return fmt.Errorf("le contact %s existe déjà", nom)
	}
	if !isValidPhoneNumber(tel) {
		return errors.New("numéro de téléphone invalide (doit contenir exactement 10 chiffres)")
	}
	id := a.getNextID()
	a.Contacts[id] = Contact{Nom: nom, Tel: tel}
	return a.Sauvegarder()
}

func (a *Annuaire) Supprimer(nom string) error {
	for id, c := range a.Contacts {
		if c.Nom == nom {
			delete(a.Contacts, id)
			return a.Sauvegarder()
		}
	}
	return fmt.Errorf("contact %s introuvable", nom)
}

func (a *Annuaire) Modifier(nom, tel string) error {
	if !isValidPhoneNumber(tel) {
		return errors.New("numéro de téléphone invalide (doit contenir exactement 10 chiffres)")
	}
	for id, c := range a.Contacts {
		if c.Nom == nom {
			a.Contacts[id] = Contact{Nom: nom, Tel: tel}
			return a.Sauvegarder()
		}
	}
	return fmt.Errorf("contact %s introuvable", nom)
}

func (a *Annuaire) Lister() []Contact {
	var liste []Contact
	for _, c := range a.Contacts {
		liste = append(liste, c)
	}
	sort.Slice(liste, func(i, j int) bool {
		return liste[i].Nom < liste[j].Nom
	})
	return liste
}
