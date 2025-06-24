# The Go Programming Language - Code Practice

Ce repository contient ma pratique du code du livre **"The Go Programming Language"** par Alan A. A. Donovan et Brian W. Kernighan.

🌐 **Livre officiel**: [https://www.gopl.io/](https://www.gopl.io/)  
📖 **Code source officiel**: [https://github.com/adonovan/gopl.io/](https://github.com/adonovan/gopl.io/)

## 📚 Structure du Repository

```
Golang_practice/
├── Chapter1_Tutorial/
│   ├── Theory/          # Exemples du chapitre
│   └── Exercise/        # Exercices pratiques
├── Chapter2_Program_Structure/
├── Chapter3_Basic_Data_Types/
├── Chapter4_Composite_Types/
├── Chapter5_Functions/
├── Chapter6_Methods/
├── Chapter7_Interfaces/
├── Chapter8_Goroutines_and_Channels/
├── Chapter9_Concurrency_with_Shared_Variables/
├── Chapter10_Packages_and_Go_Tool/
├── Chapter11_Testing/
├── Chapter12_Reflection/
├── Chapter13_Low_Level_Programming/
└── README.md
```

## 🚀 Démarrage Rapide

### Prérequis
- Go 1.19+ installé ([Installation officielle](https://golang.org/doc/install))
- Un éditeur de code (VS Code + extension Go recommandé)

### Commandes Essentielles
```bash
# Exécuter un programme Go
go run filename.go

# Compiler un programme
go build filename.go

# Exécuter tous les fichiers d'un package
go run .

# Formater le code
go fmt ./...

# Vérifier les erreurs
go vet ./...
```

## 📖 Progression par Chapitre

### ✅ Chapitre 1: Tutorial
- [x] Hello World
- [x] Arguments de ligne de commande
- [ ] Trouver des fichiers dupliqués
- [ ] GIF animé de figures de Lissajous
- [ ] Serveur web

### ⏳ Chapitre 2: Program Structure
- [ ] Noms et déclarations
- [ ] Variables
- [ ] Assignations
- [ ] Déclarations de type
- [ ] Packages et fichiers
- [ ] Portée (Scope)

### ⏳ Chapitre 3: Basic Data Types
- [ ] Entiers
- [ ] Nombres à virgule flottante
- [ ] Nombres complexes
- [ ] Booléens
- [ ] Chaînes de caractères
- [ ] Constantes

### ⏳ Chapitre 4: Composite Types
- [ ] Arrays
- [ ] Slices
- [ ] Maps
- [ ] Structs
- [ ] JSON
- [ ] Templates de texte et HTML

### ⏳ Chapitre 5: Functions
- [ ] Déclarations de fonctions
- [ ] Récursion
- [ ] Valeurs de retour multiples
- [ ] Erreurs
- [ ] Valeurs de fonction
- [ ] Fonctions anonymes
- [ ] Fonctions variadiques
- [ ] Deferred function calls
- [ ] Panic
- [ ] Recover

### ⏳ Chapitre 6: Methods
- [ ] Déclarations de méthodes
- [ ] Méthodes avec récepteurs pointeurs
- [ ] Composition d'objets par embedding
- [ ] Method values et expressions
- [ ] Exemple: bit vector type
- [ ] Encapsulation

### ⏳ Chapitre 7: Interfaces
- [ ] Interfaces comme contrats
- [ ] Types d'interface
- [ ] Satisfaction d'interface
- [ ] Parsing de flags avec flag.Value
- [ ] Interface values
- [ ] Tri avec sort.Interface
- [ ] L'interface http.Handler
- [ ] L'interface error
- [ ] Type assertions
- [ ] Discrimination de type avec type switches
- [ ] Exemple: Token-based XML decoding
- [ ] Quelques conseils

### ⏳ Chapitre 8: Goroutines and Channels
- [ ] Goroutines
- [ ] Exemple: serveur clock concurrent
- [ ] Channels
- [ ] Pipelines
- [ ] Channels unidirectionnels
- [ ] Channels avec buffer
- [ ] Looping en parallèle
- [ ] Exemple: serveur de chat concurrent
- [ ] Multiplexing avec select
- [ ] Exemple: directory traversal concurrent
- [ ] Cancellation
- [ ] Exemple: serveur de chat avec déconnexion

### ⏳ Chapitre 9: Concurrency with Shared Variables
- [ ] Race conditions
- [ ] Exclusion mutuelle: sync.Mutex
- [ ] Read/write mutexes: sync.RWMutex
- [ ] Memory synchronization
- [ ] Lazy initialization: sync.Once
- [ ] The race detector
- [ ] Exemple: banque concurrent
- [ ] Goroutines et threads

### ⏳ Chapitre 10: Packages and the Go Tool
- [ ] Introduction
- [ ] Import paths
- [ ] La déclaration package
- [ ] Import declarations
- [ ] Blank imports
- [ ] Packages et naming
- [ ] L'outil go

### ⏳ Chapitre 11: Testing
- [ ] L'outil go test
- [ ] Fonctions de test
- [ ] Coverage
- [ ] Fonctions de benchmark
- [ ] Profiling
- [ ] Fonctions d'exemple

### ⏳ Chapitre 12: Reflection
- [ ] Pourquoi la reflection?
- [ ] reflect.Type et reflect.Value
- [ ] Display, une fonction d'affichage récursive
- [ ] Exemple: encoding S-expressions
- [ ] Setting variables avec reflect.Value
- [ ] Exemple: decoding S-expressions
- [ ] Accès aux tags de struct fields
- [ ] Affichage des méthodes d'un type
- [ ] Quelques mots de prudence

### ⏳ Chapitre 13: Low-Level Programming
- [ ] unsafe.Sizeof, Alignof, et Offsetof
- [ ] unsafe.Pointer
- [ ] Exemple: deep equivalence
- [ ] Appel de code C avec cgo
- [ ] Une autre mise en garde

## 🎯 Objectifs d'Apprentissage

- [ ] Maîtriser les fondamentaux de Go
- [ ] Comprendre la programmation concurrente avec goroutines
- [ ] Apprendre les interfaces et méthodes
- [ ] Pratiquer les tests et benchmarks
- [ ] Explorer la reflection et la programmation bas niveau
- [ ] Construire des applications web simples
- [ ] Implémenter des patterns de concurrence

## 📝 Notes de Progression

Chaque chapitre contient :
- **Theory/** : Exemples tirés du livre
- **Exercise/** : Exercices à la fin de chaque chapitre

## 📚 Ressources Complémentaires

- [Site officiel du livre](https://www.gopl.io/)
- [Code source officiel](https://github.com/adonovan/gopl.io/)
- [Documentation Go](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Tour of Go](https://tour.golang.org/)
- [Go Playground](https://play.golang.org/)

## 👥 Auteurs du Livre

- **Alan A. A. Donovan** - Staff Engineer chez Google, spécialisé dans les outils de développement Go
- **Brian W. Kernighan** - Professeur à Princeton, co-auteur de "The C Programming Language"

---

*Publié le 26 octobre 2015 | Addison-Wesley | 380 pages | ISBN: 978-0134190440*
