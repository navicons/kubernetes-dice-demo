<template>
    <div id="app">
        <div class="dice-wrapper">
            <div :class="getDice"></div>
            <button @click="setDice">{{ btnTitle }}</button>
            <div v-if="resultVisible" :class="getResultClass">{{ getResultText }}</div>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
import "./assets/dark.css";
import "./assets/dice.css";

export default {
    name: "App",
    data() {
        return {
            diceMax: 6,
            diceNum: 1,
            victory: false,
            resultVisible: false,
            btnTitle: 'Roll',
            info: null,
        };
    },
    created() {
        this.setRandomDiceData();
        axios.get('/backend/settings').then(response => {
            this.diceMax = response.data.max
            this.btnTitle = response.data.btnTitle
        })
    },
    methods: {
        setRandomDiceData() {
            this.victory = false
            this.diceNum = Math.floor(Math.random() * this.diceMax) + 1;
        },
        setDice() {
            this.resultVisible = false
            axios.get('/backend/random').then(response => {
                this.diceNum = response.data.num
                this.victory = response.data.victory
                this.resultVisible = true
            })
        }
    },
    computed: {
        getDice() {
            return `dice dice-${this.diceNum}`;
        },
        getResultClass() {
            if (this.victory) {
                return `result result-winner`;
            }
            return `result result-loser`;
        },
        getResultText() {
            if (this.victory) {
                return `Winner`;
            }
            return `Loser`;
        }
    }
};
</script>

<style>
button {
    margin-top: 1em;
    margin-right: 0;
    margin-bottom: 0;
}

.dice-wrapper {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}

.result {
    margin-top: 10px;
}

.result.result-winner {
    color: green;
}

.result.result-loser {
    color: red;
}

.dice {
    font-size: 5em;
}
</style>
