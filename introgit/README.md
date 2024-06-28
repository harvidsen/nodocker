# Introgit

**git** er et verktøy som lar oss

- Ha versjonskontroll og frihet til å eksperimentere
- Samarbeide asynkront på prosjekter
- Kontrollere arbeid på en felles platform

## Hovedprinsipp

![](../drawio/git_concept.drawio.svg)

Se for deg at vi har et kodeprosjekt som består av forskjellige mapper og filer med både kode og dokumentasjon.

- Prosjektet er lagres i et sentralisert **repository** (**repo**), gjerne hos en git leverandør som for eksempel [GitHub](https://github.com), [GitLab](https://gitlab.com) eller [Bitbucket](https://bitbucket.org). Dette skal nå være "kilden til sannhet" for alle som jobber med dette prosjektet, vi kaller dette for **origin**.
- For at Håkon skal gjøre endringer i prosjektet må han først **klone** origin til min pc.
- Så kan Håkon gjøre endringer i sin lokale kopi av repoet. Når han har gjort endringer som skal være tilgjengelige for andre **push**er han tilbake til origin.
- Hvis Ole har lyst til å se og videreutvikle på dette så må han først **pull**e fra origin for å sørge for at hans lokale kopi av repoet inneholder de siste endringene.


## Commit

Hver endring som gjøres i repoet må lagres i en **commit**. En commit ser typisk slik ut

```diff
 def my_function(number):
-    """ Add 1 to input number """
-    number = number + 1
+    """ Add 2 to input number and return the square of the result"""
+    number = (number + 2)**2
     return number
```

Her endrer vi på 2 linjer i en funksjon. Git ser da at vi fjerner 2 linjer og legger til to nye.

Når vi lager flere commits bygger vi oss opp en endringshistorikk over prosjektet. Da kan vi gå tilbake til tidligere versjoner og se, eller tilbakestille prosjektet hvis vi ønsker.

## Branch og merge

![](../drawio/git_branch_merge.drawio.svg)

Hvis vi vil gjøre endringer i prosjektet og lagre i origin uten å overskrive hovedversjonen av prosjektet kan vi opprette en **branch** som kan holde på commits separat fra hovedversjonen, eller hoved-branchen. Siden commit historikken bygger opp en tilstand av prosjektet kan vi se på ulike brancher som ulike tilstander av prosjektet. Vanligvis kalles hovedbranchen for main. Et git repo sin tilstand er alltid tilknyttet en branch. Ved å bytte mellom brancher vil man se prosjektet sin tilstand på siste commit i branchen vi står på.

Når vi er fornøyde med endringene i vår nye branch kan vi **merge** en branch inn for eksempel main branch. Vanligvis gjøres dette via en **Pull Request** som andre som arbeider med prosjektet kan se på og godkjenne før vi merger.

## Overblikk

![](../drawio/git_concept_branches.drawio.svg)

Vi har nå er litt mer sammensatt bilde der commits gjøres mot én branch og vi kan pushe og pulle ting som er på utvalgt branch. Når man vil ta commits fra én branch til en annen må vi merge. Den ryddigste måten å merge på er å lage en Pull Request, godkjenne endringer sammen med en kollega og så merge.


## Nyttige tips

- Alltid jobb på din egne branch. Prøv å unngå å være flere som jobber på samme branch
- Commit små commits ofte
- Push til din egne branch
- Lag ny branch for hvert nye "tema" du arbeider med
