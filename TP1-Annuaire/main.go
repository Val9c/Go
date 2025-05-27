package main

import (
	"flag"
	"fmt"
	"log"

	"tp1-annuaire/annuaire"
)

func main() {
	action := flag.String("action", "", "Action à effectuer : ajouter, lister, supprimer, modifier, rechercher")
	nom := flag.String("nom", "", "Nom du contact")
	tel := flag.String("tel", "", "Numéro de téléphone du contact")
	flag.Parse()

	a := annuaire.NewAnnuaire()
	err := a.Charger()
	if err != nil {
		log.Fatalf("Erreur lors du chargement des contacts : %v", err)
	}

	switch *action {
	case "ajouter":
		if *nom == "" || *tel == "" {
			log.Fatal("Veuillez fournir un nom et un numéro de téléphone pour ajouter un contact.")
		}
		if err := a.Ajouter(*nom, *tel); err != nil {
			log.Println("Erreur :", err)
		} else {
			fmt.Println("Contact ajouté.")
		}

	case "lister":
		contacts := a.Lister()
		if len(contacts) == 0 {
			fmt.Println("Aucun contact enregistré.")
			return
		}
		fmt.Println("Liste des contacts :")
		for _, c := range contacts {
			fmt.Printf("- %s : %s\n", c.Nom, c.Tel)
		}

	case "supprimer":
		if *nom == "" {
			log.Fatal("Veuillez fournir un nom pour supprimer un contact.")
		}
		if err := a.Supprimer(*nom); err != nil {
			log.Println("Erreur :", err)
		} else {
			fmt.Println("Contact supprimé.")
		}

	case "modifier":
		if *nom == "" || *tel == "" {
			log.Fatal("Veuillez fournir un nom et un nouveau numéro pour modifier un contact.")
		}
		if err := a.Modifier(*nom, *tel); err != nil {
			log.Println("Erreur :", err)
		} else {
			fmt.Println("Contact modifié.")
		}

	case "rechercher":
		if *nom == "" {
			log.Fatal("Veuillez fournir un nom à vérifier.")
		}
		if a.Rechercher(*nom) {
			fmt.Printf("Le contact \"%s\" existe dans l'annuaire.\n", *nom)
		} else {
			fmt.Printf("Le contact \"%s\" n'existe pas.\n", *nom)
		}

	default:
		fmt.Println("Action non reconnue. Actions valides : ajouter, lister, supprimer, modifier, rechercher")
	}
}
