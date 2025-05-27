# TP1 - Annuaire en Go
## Description
Ce TP est un annuaire développé en Go permettant de gérer des contacts (nom + numéro de téléphone). Les contacts sont stockés dans un fichier JSON local (contacts.json).
Le programme utilise les flags de la ligne de commande pour réaliser les opérations suivantes : ajouter, lister, supprimer, modifier, rechercher.
## Équipe
Valentin ROYER
## Fonctionnalités

Ajouter un contact avec un nom et un numéro de téléphone valide (exactement 10 chiffres).

Lister tous les contacts présents dans l'annuaire.

Supprimer un contact à partir de son nom.

Modifier le numéro de téléphone d'un contact existant.

Vérifier si un contact existe déjà dans l'annuaire par son nom.

## Structure du JSON
Les contacts sont enregistrés dans un fichier JSON sous la forme :

```json
{
  "1": {
    "nom": "Charlie",
    "tel": "0811223344"
  },
  "2": {
    "nom": "Louise",
    "tel": "0612345678"
  }
}
```
Chaque contact est associé à un ID unique (clé JSON).

## Utilisation
Le programme s'utilise via des flags en ligne de commande.

Voici les flags disponibles :

| Flag     | Description                                                       | Exemple            |
|----------|-------------------------------------------------------------------|--------------------|
| --action | Action à effectuer (ajouter, lister, supprimer, modifier, rechercher) | --action ajouter   |
| --nom    | Nom du contact (requis pour toutes sauf lister)                   | --nom "Charlie"    |
| --tel    | Numéro de téléphone (requis pour ajouter et modifier)             | --tel "0811223344" |


## Exemples de commandes
Ajouter un contact :

``` bash
go run main.go --action ajouter --nom "Charlie" --tel "0811223344"
```
Lister tous les contacts :
``` bash
go run main.go --action lister
```
Supprimer un contact :
``` bash
go run main.go --action supprimer --nom "Charlie"
```
Modifier un contact :
``` bash
go run main.go --action modifier --nom "Charlie" --tel "0999888777"
```
Vérifier si un contact existe :
``` bash
go run main.go --action rechercher --nom "Charlie"
```

## Remarques
Le numéro de téléphone doit être précisément 10 chiffres. Toute autre forme sera rejetée.

Si un contact avec le même nom existe déjà, l'ajout est refusé.

Le fichier contacts.json est créé automatiquement dans le répertoire du programme lors du premier ajout.