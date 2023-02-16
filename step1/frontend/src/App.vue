<template>
  <div id="app">
    <div class="dice-wrapper">
      <div :class="getDice"></div>
      <button @click="setDice">Roll the dice!</button>
      <div v-if="resultVisible" :class="getResultClass">{{ getResultText }}</div>
    </div>
  </div>
</template>

<script>
import "./assets/dark.css";
import "./assets/dice.css";

export default {
  name: "App",
  data() {
    return {
      diceNum: 1,
      victory: false,
      resultVisible: false
    };
  },
  methods: {
    setRandomDiceData() {
      this.victory = false
      this.diceNum = Math.floor(Math.random() * 6) + 1;
      if (this.diceNum === 4) {
        this.victory = true
      }
    },
    setDice() {
      let count = 0;
      this.resultVisible = false
      const timer = setInterval(() => {
        this.setRandomDiceData();
        if (count >= 6) {
          clearInterval(timer);
          this.resultVisible = true
        }
        count += 1;
      }, 150);
    }
  },
  computed: {
    getDice() {
      this.setRandomDiceData();
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
