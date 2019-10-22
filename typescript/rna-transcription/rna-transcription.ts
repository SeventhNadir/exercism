const rnaTranslation = new Map<string, string>();

rnaTranslation.set("G", "C");
rnaTranslation.set("C", "G");
rnaTranslation.set("A", "U");
rnaTranslation.set("T", "A");


class Transcriptor {

    toRna(rna: string) {
        var translated_rna = "";
        for (var char of rna) {
            if (!rnaTranslation.get(char)) {
                throw new Error("Invalid input DNA.")
            };
            translated_rna += rnaTranslation.get(char);
        } 
        return translated_rna;
    }
}

export default Transcriptor