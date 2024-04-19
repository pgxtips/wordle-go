// Short hand for $(document).ready(function() { ... });

var _current_row_pos: number = 0;
var _current_letter_pos: number = 0;
const _working_word: string[] = ["", "", "" , "", ""];
const _valid_letters: string[] = ["A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"]


function validate_letter(letter: string, row_pos: number, letter_pos: number){
    if (!_valid_letters.includes(letter)) return false;
    if (!(row_pos < 6)) return false;
    if (!(letter_pos < 5)) return false;
    return true; 
}

function validate_enter(word: string[], row_pos: number, letter_pos: number){
    if (!(row_pos < 6)) return false;
    if (!(letter_pos === 5)) return false;

    word.forEach((letter) => {
        if (!_valid_letters.includes(letter)) return false;
    }) 

    return true; 
}

function validate_delete(row_pos: number, letter_pos: number){
    if (!(row_pos < 6)) return false;
    if (!(letter_pos <= 5 && letter_pos > 0)) return false;
    return true; 
}

function add_letter(letter: string, row_pos: number,  letter_pos: number){
    if (validate_letter(letter, row_pos, letter_pos)){
        _working_word[letter_pos] = letter;

        let id = `#row${row_pos+1}-square${letter_pos+1}`;
        $(id).html(letter)

        _current_letter_pos++;
    }
}

function delete_letter(row_pos: number,  letter_pos: number){
    if (validate_delete(row_pos, letter_pos)){
        _current_letter_pos--;
        let id = `#row${row_pos+1}-square${letter_pos}`;
        $(id).html("")
    }
}

function submit_word(word: string[], row_pos: number,  letter_pos: number){
    if (validate_enter(word, row_pos, letter_pos)){
        //check if finishing condition
        _current_row_pos++;
        _current_letter_pos = 0;
    }
}

async function test(){
    let fake_data = ["G","A","M","E","R"]
    let payload: Payload = {
        command: "submit_word",
        data: fake_data 
    }
    await call_post(payload);
}

$(function() {

    $("body").on("click", ".input-letter", function(){
        let letterEl = $(this);
        let letter: string = letterEl.html().toUpperCase();
        console.log(letter)

        let row_pos: number = _current_row_pos;
        let letter_pos: number = _current_letter_pos;
        let word: string[] = _working_word;

        if (letter == "ENTER"){
            submit_word(word, row_pos, letter_pos)
            return;
        }
        else if (letter == "DEL"){
            delete_letter(row_pos, letter_pos);
            return;
        }
        
        add_letter(letter, row_pos, letter_pos);

    })


});
