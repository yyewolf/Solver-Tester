# Wordle Solver Tester
## Pourquoi ?

Le project du deuxième semestre pour les premières années de la promotion 2024 à Telecom NANCY était le suivant : Refaire le jeu "Wordle" (WEB) ainsi qu'un Solveur pour ce jeu (en C). L'équipe nous à aussi dit qu'ils comptaient tester nos solveur sans nous donner le "comment ?".

## Protocole

Pour rappel : la stdin du solver correspond à l'entrée du logiciel (`scanf` l'utilise), et la stdout correspond à la sortie (`printf`).
Voici un exemple de partie dont la taille de la solution est 5 : 

- `stdout` : `crane` (il ne faut pas mettre d'accents)
- `stdin` : `20011`
- `stdout` : `tours`
- `stdin` : `00000`
- `stdout` : `chien`
- `stdin` : `22222`

Vous l'aurez sûrement compris :
- `0` => Faux
- `1` => Jaune
- `2` => Vert

## Faux solveur

Vous pouvez trouver dans ce repo un faux solveur (même pas en C), il sert juste à tester le testeur (et oui).

## Utilisation

Il vaut mieux utiliser ce solveur sur Linux plutôt que sur Windows.
Vous pouvez voir les arguments de la ligne de commande en faisant : `wordle_tester -h`
Pour lancer une partie simple après avoir cloner ce repo (et avoir mis votre executable dans le bon dossier) : `wordle_tester -dictionary dictionary/dict.txt -games 5 -size 5`

## Installation

Ce testeur à été fait en [Go](https://go.dev/dl/), pour compiler je vous conseilles de l'avoir.

```
git clone https://github.com/yyewolf/Solver-Tester.git 
cd solver-tester
go get .
go build
wordle_tester -dictionary dictionary/dict.txt -games 5 -size 5
```

N'oubliez pas de mettre les fichiers executables dans le bon dossier !!