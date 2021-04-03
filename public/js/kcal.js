const ACTIVITY = {
    '0': 1.0,
    '1': 1.1,
    '2': 1.2,
    '3': 1.3,
    '4': 1.4,
    '5': 1.5,
    '6': 1.6,
    '7': 1.7
}
const GOAL = {
    '0': -500.0,
    '1': 0.0,
    '2': 500.0
}

function calculate() {
    var gender = document.getElementById('gender').value;
    var age = document.getElementById('age').value;
    var height = document.getElementById('height').value;
    var weight = document.getElementById('weight').value;
    var activity = document.getElementById('activity').value;
    var goal = document.getElementById('goal').value;

    var score = 0.0;
    switch (gender) {
        case "women":
            score = (655.1 + (9.563 * weight) + (1.85 * height) - (4.676 * age)) * ACTIVITY[activity] + GOAL[goal];
            break;
        case "men":
            score = (66.5 + (13.75 * weight) + (5.003 * height) - (6.775 * age)) * ACTIVITY[activity] + GOAL[goal];
            break;
        default:
            score = 10;
    }

    score = Math.round(score)

    renderCalculateOutput(score)
}

function renderCalculateOutput(score) {
    var tag = document.createElement('p');
    tag.appendChild(document.createTextNode('Wynik: ' + score));

    score = document.getElementById('score');
    if (score.hasChildNodes()){
        while(score.lastElementChild) {
            score.removeChild(score.lastElementChild)
        }
    }
    document.getElementById('score').appendChild(tag)
}