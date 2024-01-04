# advent_of_code_2023

2023day1 : 
La principale difficulté : comprendre comment convertir les bytes en chaînes de caractères puis en int.
Grâce au package strconv, se fait finalement.
J'ai fait des fonctions auxiliaires pour savoir si un indice dans une chaîne de caractères pouvait être le début ou la fin d'un chiffre en toutes lettres. Pour ce faire, j'ai boulé sur une liste ["one", "two"...]

2023day2 :
Beaucoup de parsing : on utilise la fonction strings.Split().
J'ai créé une structure Set et j'ai fait 2 fonctions de parsing : la première pour les lignes et la seconde pour séparer les éléments de la liste [3 blue, 2 green] par exemple.
Au début, je n'avais pas parsé autant mais ça m'a été nécessaire pour la partie 2.

2023day4 :
Ce jour ressemblait au day2 : du parsing assez précis et la création de structures (j'ai créé une structure Card, et les méthodes associées). Le résultat final est directement donné par des méthodes sur le type Card. La seule fonction (autre que celle qui lit le fichier) qui n'est pas une méthode, est la fonction de parsing qui crée l'array de Cards sur lequel les calculs seront faits.

2023day10 : 
Un jour plus difficile que les premiers : je n'ai pas réussi à obtenir le résultat souhaité en partie 1.
Comme pour les jours précédents, j'avais d'abord fait une structure Tile, mais elle m'a semblée inutile à terme, et j'ai remplacé mon array de Tile par un type [][]string
La difficulté était aussi dans la taille de l'input, qui dissuadait de faire un algorithme naïf. J'essayais donc de faire un loop, mais sans succès. J'ai cherché à m'aider des solutions sur internet, mais j'avais du mal à bien comprendre leur principe et donc à les débugger efficacement.

2023day13 :
Jour difficile aussi : d'abord une fonction de parsing assez faisable, mais ensuite des fonctions de détections des réflexions très complexes et sujettes aux erreurs. J'ai tout de même pu finir la partie 1. Pour la partie 2, je n'ai pas trouvé d'autre moyens que de faire une double boucle sur i et j, soit de la complexité en n^2 au moins.

2023day17 :
Jour très difficile pour moi : je n'ai pas réussi à obtenir le bon résultat pour la partie 1, après 2 tentatives différentes (j'ai laissé les 2 fichiers sur le git). D'abord, une fonction de parsing assez faisable avec l'habitude et la maîtrise des fonctions utiles du package strconv.
Ensuite, à cause de la grande taille de l'input, j'ai esssayé d'utiliser des arrays que je traitais comme des queues de priorité. Dans ma première tentative (fichier first attempt), j'avais une fonction Calculates_path extrêmement compliquée, qui utilisait une file d'attente en la remplissant et en faisant pop() jusqu'à ce qu'elle soit vide. Cette fonction était tellement compliquée que je n'ai pas réussi à la débugger pour aboutir.
Ensuite, pour ma 2ème tentative, je me suis inspirée des solutions en ligne et j'ai créé beaucoup de fonctions auxiliaires (que j'ai bien testées) pour recoder les fonctions des files de priorité. Cependant, ma fonction principale (beaucoup plus simple que dans ma première tentative) était toujours assez obscure et je n'ai pas réussi à la débugger pour qu'elle donne le bon résultat.

2023day18 :
Encore un jour trop difficile pour moi. Pour ce qui est du parsing, pas de difficulté. Ensuite, j'ai créé une fonction très compliquée (mais testée et fonctionnelle) pour resize mon tableau de Tiles. J'avais dans l'idée que mon tableau serait redimensionné à chaque instruction faisant sortir du tableau existant (initialisé à une seule tile). Ensuite, j'ai fait une fonction pour délimiter la frontière du gouffre creusé : c'est cette fonction qui m'a posé problème. Mon traitement des instructions ne se faisait pas de manière assez fluide, et j'ai passé beaucoup de temps à essayer de débugger cette fonction compliquée. Ensuite, j'ai fait un algorithme de flood fill (avec l'aide du site red blob games entre autres), qui est compliqué mais qui ne pose pas problème en premier lieu. Je n'ai finalement pas réussi la partie 1.

2023day20 :
Parsing un peu compliqué mais faisable sans problème avec l'habitude. J'ai fait une fonction de parsing en essayant de réduire la complexité. Pour ce faire, j'indiquais les modules input et outputs par leur indice dans la liste de modules, au lieu d'aller les modifier un par un dans la liste de modules. Ensuite, j'ai fait une petite fonction par type de module pour simplifier le code final, sans trop de problèmes. Cependant, ma fonction finale Processes_pulses ne fonctionnait pas pour le 2ème exemple donné. Même après débuggage, je n'ai pas réussi à comprendre pourquoi. J'ai quand même un fichier qui marche pour le premier exemple, mais je n'ai pas le bon résultat pour la partie 1.


En résumé, l'advent of code m'a pris énormément de temps, parce que les jours étaient difficiles (sauf la première semaine, qui était faisable en y mettant le temps) et que je m'obstinais à débugger longuement sans obtenir de résultat. Cependant, j'ai vu que j'avais beaucoup progressé sur la lecture de fichier, le parsing (même complexe) et les structures (parfois très complexes aussi) en go. C'est sur les fonctions les plus compliquées, avec des récursivités par exemple que je faisais des erreurs que je ne savais pas débugger.