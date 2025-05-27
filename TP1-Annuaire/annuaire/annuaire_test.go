package annuaire

import (
	"os"
	"testing"
)

func setupTestAnnuaire() *Annuaire {
	_ = os.Remove(fichier)

	a := NewAnnuaire()
	return a
}

func TestAjouterEtRecherche(t *testing.T) {
	a := setupTestAnnuaire()

	err := a.Ajouter("Alice", "0601020304")
	if err != nil {
		t.Errorf("Erreur lors de l'ajout : %v", err)
	}

	if !a.Rechercher("Alice") {
		t.Error("Le contact Alice aurait dû exister")
	}

	err = a.Ajouter("Alice", "0600000000")
	if err == nil {
		t.Error("Erreur attendue : tentative d'ajouter un contact déjà existant")
	}
}

func TestLister(t *testing.T) {
	a := setupTestAnnuaire()

	_ = a.Ajouter("Charlie", "0707070707")
	_ = a.Ajouter("Bob", "0808080808")

	liste := a.Lister()
	if len(liste) != 2 {
		t.Errorf("Il devrait y avoir 2 contacts, obtenu : %d", len(liste))
	}

	if liste[0].Nom != "Bob" {
		t.Errorf("Le premier contact devrait être Bob, obtenu : %s", liste[0].Nom)
	}
}

func TestModifier(t *testing.T) {
	a := setupTestAnnuaire()

	_ = a.Ajouter("David", "0909090909")

	err := a.Modifier("David", "0999999999")
	if err != nil {
		t.Errorf("Erreur lors de la modification : %v", err)
	}

	if !a.Rechercher("David") || a.Contacts["1"].Tel != "0999999999" {
		found := false
		for _, c := range a.Contacts {
			if c.Nom == "David" && c.Tel == "0999999999" {
				found = true
				break
			}
		}
		if !found {
			t.Error("Le numéro de David n'a pas été modifié correctement")
		}
	}
}

func TestSupprimer(t *testing.T) {
	a := setupTestAnnuaire()

	_ = a.Ajouter("Eve", "0123456789")
	err := a.Supprimer("Eve")
	if err != nil {
		t.Errorf("Erreur lors de la suppression : %v", err)
	}

	if a.Rechercher("Eve") {
		t.Error("Le contact Eve aurait dû être supprimé")
	}

	err = a.Supprimer("Eve")
	if err == nil {
		t.Error("Erreur attendue lors de la suppression d'un contact inexistant")
	}
}
