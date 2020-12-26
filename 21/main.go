package main

import (
	"github.com/oriolf/adventofcode2020/util"
	"sort"
	"strings"
)

type meal struct {
	ingredients []string
	allergens   []string
}

type ingredient struct {
	name     string
	allergen string
}

func main() {
	util.Solve(solve1, solve2)
}

func solve1(lines []string) interface{} {
	meals, allergensMap := parseInput(lines)
	nonAllergensIngredients := getNonAllergensIngredients(meals, allergensMap)

	var count int
	for _, m := range meals {
		for _, ing := range m.ingredients {
			if _, ok := nonAllergensIngredients[ing]; ok {
				count++
			}
		}
	}

	return count
}

func solve2(lines []string) interface{} {
	_, allergensMap := parseInput(lines)
	for !allAllergensKnown(allergensMap) {
		allergensMap = discoverAllergen(allergensMap)
	}

	var ingredients []ingredient
	for k, v := range allergensMap {
		ingredients = append(ingredients, ingredient{name: v[0], allergen: k})
	}

	sort.Slice(ingredients, func(i, j int) bool {
		return ingredients[i].allergen < ingredients[j].allergen
	})

	var l []string
	for _, ing := range ingredients {
		l = append(l, ing.name)
	}

	return strings.Join(l, ",")
}

func getNonAllergensIngredients(meals []meal, allergensMap map[string][]string) map[string]struct{} {
	nonAllergens := make(map[string]struct{})
	for _, m := range meals {
	LOOP:
		for _, ing := range m.ingredients {
			for _, ingredients := range allergensMap {
				if util.StringInSlice(ing, ingredients) {
					continue LOOP
				}
			}
			nonAllergens[ing] = struct{}{}
		}
	}

	return nonAllergens
}

func allAllergensKnown(m map[string][]string) bool {
	for _, v := range m {
		if len(v) > 1 {
			return false
		}
	}
	return true
}

func discoverAllergen(m map[string][]string) map[string][]string {
	var toRemove []string
	for _, v := range m {
		if len(v) == 1 {
			toRemove = append(toRemove, v[0])
		}
	}

	for k, v := range m {
		if len(v) > 1 {
			var newv []string
			for _, x := range v {
				if !util.StringInSlice(x, toRemove) {
					newv = append(newv, x)
				}
			}
			m[k] = newv
		}
	}

	return m
}

func parseInput(lines []string) (meals []meal, m map[string][]string) {
	for _, l := range lines {
		meals = append(meals, parseMeal(l))
	}

	allergensMap := make(map[string][]string)
	for _, m := range meals {
		for _, a := range m.allergens {
			if _, ok := allergensMap[a]; !ok {
				allergensMap[a] = m.ingredients
			} else {
				allergensMap[a] = util.IntersectStrings(m.ingredients, allergensMap[a])
			}
		}
	}

	return meals, allergensMap
}

func parseMeal(l string) meal {
	parts := strings.Split(l, " (contains ")
	ingredients := parts[0]
	allergens := strings.Trim(parts[1], ")")
	return meal{ingredients: strings.Fields(ingredients), allergens: strings.Split(allergens, ", ")}
}
