<template>
    <div class="flex flex-col gap-2">
        <label :class="['block text-xs font-light', getLabelTextColor]"> {{ textAreaLabel }} </label>
        <label :class="['relative block rounded-md border shadow-sm', getHoverClassInput, getBorderClassInput]">
            <textarea
                :class="['text-xs w-full font-light p-2 peer border-none bg-transparent placeholder-transparent focus:border-transparent focus:outline-none focus:ring-0', getLabelTextColor]"
                v-model="text">
        </textarea>
        </label>
        <span v-show="showError" class="text-xs font-light text-red-400 pointer-events-none">
            {{ textAreaErrorLabel }}
        </span>
    </div>
</template>

<script>
export default {
    name: 'InputComponent',
    props: {
        textAreaLabel: {
            type: String,
            default: ''
        },
        textAreaErrorLabel: {
            type: String,
            default: 'This Field cannot be empty'
        },
        value: {
            type: String,
            default: ''
        },
        showError: {
            type: Boolean,
            default: false
        }
    },
    data() {
        return {
            text: this.value
        };
    },
    watch: {
        value(newVal) {
            this.text = newVal;
            this.checkEmpty();
        },
        text(newVal) {
            this.$emit('update:value', newVal);
            this.checkEmpty();
        }
    },
    computed: {
        getHoverClassInput() {
            return this.showError ?
                "focus-within:border-red-400 focus-within:ring-1 focus-within:ring-red-400" :
                "focus-within:border-blue-500 focus-within:ring-1 focus-within:ring-blue-500";
        },
        getBorderClassInput() {
            return this.showError ? "border-red-400" : "border-gray-200";
        },
        getLabelTextColor() {
            return this.showError ? "text-red-400" : "text-gray-700";
        }
    },
    methods: {
        checkEmpty() {
            if (this.text.trim() !== '') {
                this.$emit('update:showError', false);
            }
        }
    }
};
</script>