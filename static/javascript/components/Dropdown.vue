<template>
    <div @focusout="handleFocusOut" class="flex flex-col gap-2" tabindex="0">
        <label v-show="showDropdownLabel" :class="['block text-xs font-light', getLabelTextColor]"> {{ dropdownLabel }} </label>
        <div class="dropdown relative">
            <div class="relative block">
                <div :class="['inline-flex items-center overflow-hidden rounded-md border bg-white w-full', getBorderClassInput ]">
                    <div @click="toggleDropdown"
                        :class="['border-e px-4 py-2 text-xs font-light hover:bg-gray-50 w-full', getLabelTextColor, getBorderClassInput]">
                        {{ selectedValue }}
                    </div>
                    <button @click="toggleDropdown"
                        :class="['h-full p-2 hover:bg-gray-50 hover:text-gray-700',getLabelTextColor]">
                        <span class="sr-only">Menu</span>
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 20 20" fill="currentColor">
                            <path fill-rule="evenodd"
                                d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"
                                clip-rule="evenodd" />
                        </svg>
                    </button>
                </div>
            </div>
            <div :class="['absolute left-0 end-0 z-10 mt-2 divide-y divide-gray-100 rounded-md border border-gray-100 bg-white shadow-lg w-full', showDropdown]"
                role="menu">
                <div class="p-2">
                    <div v-for="item in dropdownItems" :key="item.value" :data-value="item.value"
                        class="block rounded-lg px-4 py-2 text-xs font-light hover:bg-gray-50 hover:text-gray-700"
                        role="menuitem" @click="handleSelected(item)">
                        {{ item.key }}
                    </div>
                </div>
            </div>
        </div>
        <label v-show="showError" :class="['block text-xs font-light pb-2', getLabelTextColor]"> Please select item from dropdown </label>

    </div>
</template>

<script>
export default {
    name: 'Dropdown',
    props: {
        value: {
            type: String,
            default: ''
        },
        showError: {
            type: Boolean,
            default: false
        },
        dropdownLabel: {
            type: String,
            default: ''
        },
        dropdownItems: {
            type: Array,
            default: () => [
            ]
        }
    },
    data() {
        return {
            dropdownOpen: false,
            selectedItem: 'Please choose',
        }
    },
    computed: {
        showDropdown() {
            return this.dropdownOpen ? '' : 'hidden';
        },
        selectedValue() {
            const selectedItem = this.dropdownItems.find(item => item.value === this.value);
            return selectedItem ? selectedItem.key : this.selectedItem;
        },
        showDropdownLabel() {
            return this.dropdownLabel !== '' ? true : false;
        },
        getBorderClassInput() {
            return this.showError ? "border-red-400" : "border-gray-200";
        },
        getLabelTextColor() {
            return this.showError ? "text-red-400" : "text-gray-700";
        }
    },
    methods: {
        toggleDropdown() {
            this.dropdownOpen = !this.dropdownOpen;
        },
        handleFocusOut(event) {
            if (!this.$el.contains(event.relatedTarget)) {
                this.dropdownOpen = false;
            }
        },
        handleSelected(item) {
            this.dropdownOpen = false;
            this.selectedItem = item.key;
            this.$emit('update:value', item.value);
            this.$emit('update:show-error', false);
        }
    }
};
</script>