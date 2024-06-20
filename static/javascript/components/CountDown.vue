<template>
    <section class="count-down-timer">
        <div class="container mx-auto w-full-h-full">
            <div :class="positionClass">
                <div class="flex flex-col justify-center text-center">
                    <span class="text-2xl countdown">
                        <span>{{ days }}</span>
                    </span>
                    days
                </div>
                <div class="flex flex-col justify-center text-center">
                    <span class="text-2xl countdown">
                        <span>{{ hours }}</span>
                    </span>
                    hours
                </div>
                <div class="flex flex-col justify-center text-center">
                    <span class="text-2xl countdown">
                        <span>{{ minutes }}</span>
                    </span>
                    min
                </div>
                <div class="flex flex-col justify-center text-center">
                    <span class="text-2xl countdown">
                        <span>{{ seconds }}</span>
                    </span>
                    sec
                </div>
            </div>
        </div>
    </section>
</template>

<script>
import moment from 'moment';
import 'moment/locale/id';

moment.locale('id');

export default {
    props: {
        targetDate: {
            type: String,
            required: true
        },
        position: {
            type: String,
            default: 'center'
        }
    },
    data() {
        return {
            days: 0,
            hours: 0,
            minutes: 0,
            seconds: 0,
            countDown: 0,
            intervalId: null
        };
    },
    computed: {
        positionClass() {
            return this.position === 'center'
                ? 'flex justify-center gap-5'
                : this.position === 'left'
                    ? 'flex gap-5 justify-start'
                    : 'flex gap-5 justify-end';
        }
    },
    mounted() {
        const countDownDate = moment(this.targetDate, "YYYY-MM-DD").valueOf();
        this.intervalId = setInterval(() => {
            this.countDown = countDownDate - new Date().getTime();
            this.updateTime();
        }, 1000);
    },
    beforeDestroy() {
        clearInterval(this.intervalId);
    },
    methods: {
        updateTime() {
            this.days = Math.floor(this.countDown / (1000 * 60 * 60 * 24));
            this.hours = Math.floor((this.countDown % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
            this.minutes = Math.floor((this.countDown % (1000 * 60 * 60)) / (1000 * 60));
            this.seconds = Math.floor((this.countDown % (1000 * 60)) / 1000);
            if (this.days < 0) this.days = 0;
            if (this.hours < 0) this.hours = 0;
            if (this.minutes < 0) this.minutes = 0;
            if (this.seconds < 0) this.seconds = 0;
        }
    }
};
</script>