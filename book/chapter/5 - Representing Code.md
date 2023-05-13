A _regular language_ is the rules for how characters get grouped into tokens.

A formal grammar takes a set of atomic pieces it calls its "alphabet".

| Terminology         | Lexical grammar  | Syntactic grammar |
| ------------------- | ---------------- | ----------------- |
| Alphabet            | Characters       | Tokens            |
| Words/String        | Tokens or Lexeme | Expressions       |
| It's implemented by | Scanner          | Parser            |

A formal grammar's job is to take a string of characters and turn it into a tree of nested elements.

Strings derived from the tules of the grammar are called _derivations_. Rules are called _productions_ because they produce string in the grammar.

Each production in a context-free grammar has a **head** and a **body**. The head is its name. The body is a sequence symbols.

- A _terminal_ is a symbol that can't be broken down any further.
- A _nonterminal_ is a symbol that can be broken down further. It's a name reference to another rule in the grammar.

There is one last refinement: you may have multiple rules with the same name. When you reach a nonterminal with that name, you are allowed to pick any of the rules for it, whichever floats your boat.

Each rule is a name, followed by an arrow, followed by a sequence of symbols, and finally a semicolon.

```
rule-name -> symbol1 symbol2 symbol3;

breakfast  → protein "with" breakfast "on the side" ;
breakfast  → protein ;
breakfast  → bread ;

protein    → crispiness "crispy" "bacon" ;
protein    → "sausage" ;
protein    → cooked "eggs" ;

crispiness → "really" ;
crispiness → "really" crispiness ;

cooked     → "scrambled" ;
cooked     → "poached" ;
cooked     → "fried" ;

bread      → "toast" ;
bread      → "biscuits" ;
bread      → "English muffin" ;
```

We can use the grammar to generate random breakfasts. Let's play a round and see how it works. By age-old convention, the game starts with the first rule in the grammar, here `breakfast`. There are three productions for that, and we randomly pick the first one. Our resulting string looks like:

```
protein "with" breakfast "on the side"
```

We pick the first production for `protein` and get:

```
protein → cooked "eggs" ;
"poached" "eggs" "with" breakfast "on the side"
```

Any time we hit a rule that had multiple productions, we just picked one arbitrarily.

### Enhancing our notation

1. Instead of repeating the rule name each time we want to add another production for it, we'll allow a series of productions separeed by a pipe character (`|`).

```
bread → "toast" | "biscuits" | "English muffin" ;
```

2. Further, we'll allow parentheses for grouping and then allow | within that to select on from a series of options within the middle of a production.

```
protein → ( "scrambled" | "poached" | "fried" ) "eggs" ;
```

3. Using recursion to support repeated sequences of symbols has a certain appealing purity, but it's kind of a chore to write. We'll allow a `*` after a symbol to mean "zero or more of these".

```
crispiness → "really" "really"* ;
```

4. We'll allow `+` to mean "one or more of these".

```
crispiness → "really"+ ;
```

5. We'll allow `?` to mean "zero or one of these".

```
breakfast → protein ( "with" breakfast "on the side" )? ;
```

With these changes, our grammar looks like this:

```
breakfast → protein ( "with" breakfast "on the side" )?
          | bread ;

protein   → "really"+ "crispy" "bacon"
          | "sausage"
          | ( "scrambled" | "poached" | "fried" ) "eggs" ;

bread     → "toast" | "biscuits" | "English muffin" ;
```


```
expression     → literal
               | unary
               | binary
               | grouping ;

literal        → NUMBER | STRING | "true" | "false" | "nil" ;
grouping       → "(" expression ")" ;
unary          → ( "-" | "!" ) expression ;
binary         → expression operator expression ;
operator       → "==" | "!=" | "<" | "<=" | ">" | ">="
               | "+"  | "-"  | "*" | "/" ;
```

In addition to quoted strings for terminals that match exact lexemes, we CAPITALIZE terminals that are single lexemes whose text representation may vary.

