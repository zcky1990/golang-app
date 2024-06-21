<template>
    <span class="flex items-center">
        <span v-if="isLeft" :class="['shrink-0 px-6 text-sm font-light', getTextColor, messageDecoration]"
            v-show="hasMessage">
            {{ message }}
        </span>
        <span :class="['h-px flex-1', getColor]"></span>
        <span v-if="isCenter" :class="['shrink-0 px-6 text-sm font-light', getTextColor, messageDecoration]"
            v-show="hasMessage">
            {{ message }}
        </span>
        <span :class="['h-px flex-1', getColor]"></span>
        <span v-if="isRight" :class="['shrink-0 px-6 text-sm font-light', getTextColor, messageDecoration]"
            v-show="hasMessage">
            {{ message }}
        </span>
    </span>
</template>

<script>
export default {
    props: {
        message: {
            type: String,
            default: ''
        },
        messagePosition: {
            type: String,
            default: 'center'
        },
        messageTransform: {
            type: String,
            default: 'uppercase'
        },
        color: {
            type: String,
            default: 'gray'
        }
    },
    computed: {
        hasMessage() {
            return this.message !== '';
        },
        isCenter() {
            return this.messagePosition === 'center';
        },
        isLeft() {
            return this.messagePosition === 'left';
        },
        isRight() {
            return this.messagePosition === 'right';
        },
        getColor() {
            return {
                gray: 'bg-gray-400',
                black: 'bg-black',
                red: 'bg-red-400'
            }[this.color] || 'bg-black';
        },
        getTextColor() {
            return {
                gray: 'text-gray-400',
                black: 'text-black',
                red: 'text-red-400'
            }[this.color] || 'text-black';
        },
        messageDecoration() {
            return {
                uppercase: 'uppercase',
                lowercase: 'lowercase',
                capitalize: 'capitalize',
                'normal-case': 'normal-case'
            }[this.messageTransform] || 'normal-case';
        }
    }
};
</script>
