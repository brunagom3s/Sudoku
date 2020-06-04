<template>
    <div class="content">
        <h1>{{ title }}</h1>
        
        <div class="grid" >
            <div class="row" v-for="(row, rowIndex) in sudoku" :key="rowIndex"> 
                <div class="cell" :class="{'border-right': colIndex === 2 || colIndex === 5, 
                                           'border-bottom': rowIndex === 2 || rowIndex === 5, 
                                           'original': cell.original,
                                           'active': activeRow === rowIndex && activeCol === colIndex, }"
                v-for="(cell, colIndex) in row" :key="colIndex"
                @click="setCellActive(rowIndex, colIndex, cell.original)">
                    <input type="number" min="1" max="9" maxlength="3" v-if="cell.value === null">
                    {{ cell.value }}
                </div>
            </div>
        </div>
        
        <button @click="getSudoku">
            Nueva partida
        </button>

        
        <!-- TIMER -->
        <div class="timer">
            <span class="elapsed" :class="{blinking : timer == null}">
                <span v-if="hours>0">{{ hours }}:</span>{{ minutes }}:{{ seconds }}
            </span>

            <!-- Start Timer -->
            <button class="timer-btns" id="start" v-if="!timer" @click="startTimer">
                <img src="../assets/images/play-btn.png"> 
            </button>
            
            <!-- Pause Timer -->
            <button class="timer-btns" id="stop" v-if="timer" @click="stopTimer">
                <img src="../assets/images/pause.png">
            </button>

        </div>   
    </div>
        

   

</template>

<script>

    export default {
        name: "Sudoku",
        data() {
            return {
                title: "sudoku",
                message: " ",
                sudoku: [],
                activeRow: -1,
                activeCol: -1,
                timer: null,
                totalTime: 0,
            };
        },
        methods: {
            getSudoku: function() {
                var self = this;
                window.backend.generateSudoku().then(result => {
                    self.sudoku = result.map(row => {
                        return row.map(cell => {
                            return {
                                value: cell !== 0 ? cell : null,
                                original: cell !== 0
                            }
                        })
                    });
                });
            },
            setCellActive(row, col, original) {
                if (original) {
                    return
                }
                if(this.activeRow === row && this.activeCol === col) {
                    this.activeRow = -1
                    this.activeCol = -1
                    return
                }
                this.activeRow = row
                this.activeCol = col
            },

            startTimer: function() {
                if (!this.timer) {
                this.timer = setInterval(() => this.countdown(), 1000);
                }
            },
            countdown: function() {
                this.totalTime++;
            },
            stopTimer: function() {
                clearInterval(this.timer);
                this.timer = null;
                this.resetButton = true;
            },
            padTime: function(time) {
                return (time < 10 ? '0' : '') + time;
            },    
        },
        computed: {
            hours: function() {
                const hours = Math.floor(this.totalTime / (60*60))
                return hours;
            },          
            minutes: function() {
                const minutes = Math.floor(this.totalTime / 60) - (this.hours * 60);
                return this.padTime(minutes);
            },
            seconds: function() {
                const seconds = this.totalTime - (this.minutes * 60) - (this.hours *60*60);
                return this.padTime(seconds);   
            }
        }   
    }    
</script>

<style scoped>
    
    /* sudoku */
    .grid {
        width: calc(9*40px);
        margin: 0.5rem, auto;
    }

    .row {
        display: flex;
        align-items: center;
        justify-content: space-between;
    }

    .cell {
        display: block;
        width: 40px;
        height: 40px;
        box-sizing: border-box;
        border: 1px solid #888;

        line-height: 40px;
        text-align: center;
    }

    .cell.border-right {
        border-right-width: 5px;    
    }
    .cell.border-bottom {
        border-bottom-width: 5px;    
    }
    .cell.original {
        font-weight: bold;
    }
    .cell:not(.original) {
        cursor: pointer;
    }
    .cell.active {
        background-color: #888;
        color: black
    }

    input {
        width: 32px;
        height: 35px;
        text-align: center;
        font-size: 20px;
        padding: 0;
        border: 3px transparent solid;
        background-color: transparent;
        color: #eee;
    }

    /* timer */
    .timer {
        font-size: 2em;
        color: #ccc;
        text-align: center;
        padding:15px;
        margin:0 auto;
    }

    .timer-btns {
        border:none;
        color: #ccc;
        font-size:1em;
        background-color: transparent;
        cursor: pointer;
    }:hover {
        color: #888;
    } 
    .timer-btns img {
        width: 1em;
        height: 1em;
    
    }

</style>