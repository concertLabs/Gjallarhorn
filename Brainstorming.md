= Gjallarhorn. =

== Allgemein ==
Gjallarhon (altnordisch "laut tönendes Horn") ist eine in Go geschriebene
Software, die zur Verwaltung einer musikalischen Notensammlung dient. In erster
Linie soll dieses Projekt nicht nur den einfachen Verwaltungsrahmen abdecken,
den man sicher ohne weiteres in z.B. Excel abbilden könnte, vielmehr soll versucht
werden das komplette Orchester digital abzubilden.


== Erfassung ==

Dabei wird das einzelne Lied eines Komponisten im digitalen Format eingescannt
und im Anschluss mit entsprechenden Metadaten in einer Datenbank abgelegt. Dabei
werden gleichzeitig die Noten des Notenblattes digital erkannt und die einfache
Stimme in ein digitales OpenMusicXML umgewandelt, dadurch steht später die
Möglichkeit offen, aus jeder Stimme eine Audio-Datei zu generieren.

== Verwaltung == 
Desweiteren soll die Datenbank dem Dirigenten die Möglichkeit bieten ein
Repertoire zusammenzustellen und diese chronologisch zu bewahren. Zudem soll er
in der Lage sein für einzelne Auftritte angepasste Listen erstellen zu können,
um auf entsprechende Umstände vorbereitet zu sein (Kirmes, Prozession o.ä).

Da der Vorgang des Erfassens im Idealfall automatisch abläuft, soll man eben-
falls in der Lage die Qualität der Notenblätter zu bewerten und zu bearbeiten,
denn es kann häufig vorkommen, dass eingescannte Noten älter als das Internet
sind. Dafür soll eine Möglichkeit geboten werden, diese mangelhaften Noten ent-
eder zu löschen oder nachträglich zu bearbeiten.

== Verteilung ==
Das System ist zudem in er Lage die Daten auf Clienten zu übertragen. Dabei soll
es keine Rolle spielen, ob es entfernte Desktop-PCs sind oder Smartphones. Das
System bietet jedem Musiker die Möglichkeit die Notendatenbank zu durchstöbern
und Noten anzusehen.

== Auftritt-Modus ==
Der Auftritt-Modus sieht vor, dass alle Musiker statt der Notenmappe ein Tablet
verwenden. Die Dirigent nutzt ebenfalls ein Tablett und dient dabei als
quasi-Server. sowohl server als auch Client haben eine App, die regelt welches
Lied momentan aufliegt. Der Dirigent dient dabei als Server und die Musiker als
Clients. Die Technik dahinter würde auf normalem WLAN basieren.


....

== Key Features == 
    * automatisches Einsortieren der eingescannten Noten anhand $Parameter.
    - Sotierung der Noten anhand eindeutiger Muster (Komponist, Name, Genre,
      $Lagerort o.ä)
    - Erstellung individueller Liederlisten für Auftritte, ständige Notenmappe,
      Ständchen o.ä.
    - Bewertung eingescannter PDFDateien und deren OCR-Pendants zur Verbesserung
      derr
    - Bewahrung des musikalischen Erbes.
    - (sobald Copyrightechnisch geklört (was halt auch Schmarrn ist imho)) Aus-
      tausch des Liedguts unterhalbt der Musikvereine.
    - Direkter Zugriff auf die Notendatenbank durch eine Android-App
    - Zwischenserver, der für Auftritte PDF-Dateien vorhält und verteilt.
    - Der Dirigent entscheidet dabei welches Lied gespielt wird und "verteilt"
      das Lied auf alle Tablets der Musiker.


== Definitionen ==
Um Unklarheiten aus dem Weg zu gehen folgen einige Definitionen, die im Projekt
weiterhin verwendet werden.

=== Noten ====
=== Stück ====
=== Lied ===
=== Instrument ===
=== Stimme ===
=== Repertoire ===
=== Sammlung ===
=== Auswahl ====



== Notenlagerung ===

Die ''Notenlagerung'' (passendes Wort finden) beschreibt nur den physischen
Aufentahltsort der Noten. Das Lagerungsystem ansich kann aber unterteilt werden
in folgendem System:
    - Lagerort (räumlicher Ort), zb Strasse + Hausnummer
    - Lagertyp (Schrank, Regal o.ä)
    - Schublade
    - Reihe in Schublade

Man sollte beachten, dass die Deklaration der Notenlagerung _nicht_ global
definieren kann, vielmehr sollte man pro Lagerort ein eindeutiges System ver-
wenden, mit dem man später Noten eindeutig finden kann.



UI:
evtl. was Racktable artiges

Datenspeicherung:
Noten nach folgendem Format ablegen: $komponist/$stück/$stimme.pdf //Phils preference
oder
Noten direkt in die DB
Als Datenbank couchdb

Storageverwaltung:
lagerort 1 : n schränke 1:n schubladen 1:n sätzen 





//// Random foo
- Eigenschaften der Titel: 
    - Seitenverhältnis (Ich habe es jetzt öfters gesehen, dass es keine standard
    A4, A3 oder A5 Blätter waren, sondern exotische xy Formate
    - Fremdschlüssel mit Notiz welche anderen Titel auf dem selben Notenblatt 
    vorhanden sind, sei es auf der Rückseite oder auf der untern Hälfte auf dem
    selben Blatt (evtl slice)

- evtl ein kleines Wiki um verschiedene Sachen zu notieren, die für einen Noten-
wart wichtig sind. ZB. Seitenverhältnisse von A5->A4 und umgekehrt, um Kopien
auf andere Formate zu kopieren.
