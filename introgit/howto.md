# How to git

Her viser vi noen nyttige kommandoer med git kommandolinjeverktøy, som lar oss gjøre der som står i [git introduksjonen](introgit/README.md). Dette kommandolinjeverktøyet er en av mange måter å arbeide med git på. Uansett hva man ønsker å bruke så er konseptene de samme og terminologien lik. Det viktigste er å finne et verktøy som gjør at man får oversikt og effektivt kan løse sine oppgaver.

Gjerne prøv ut dette på dette repoet eller ditt eget.

## Kommandoer

**clone:** Kopierer origin repo til egen lokal kopi
```bash
git clone https://github.com/harvidsen/teach.git
```

**checkout:** Bytt gren vi står på
```bash
git checkout some-branch-name
```
eller hvis branch ikke allerede eksisterer og vi vil opprette en ny
```bash
git checkout -b some-branch-name
```

**pull:** Henter ned endringer i branch til lokal kopi
```bash
git pull
```

**push:** Dytter opp endringer i branch fra lokal kopi til origin
```bash
git push
```

**status:** Se hva git har fått med seg
```
git status
```

**add:** La git se endringer du har gjort i en fil
```bash
git add ./pathto/filetoadd.txt
```
eventuelt for å adde alle filer under en mappe
```bash
git add .
```

**commit:** Lage endringsmelding av filendringer som allerede er **staget** med `git add` til branchen vi står på
```bash
git commit
```
eller for å skrive commit melding direkte
```bash
git commit -m "Some short concise message"
```

