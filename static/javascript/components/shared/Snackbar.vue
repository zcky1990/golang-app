<template>
    <div v-if="show" :class="['fixed', positionClasses, 'p-4', bgColor, 'text-white']">
      {{ message }}
      <button @click="closeSnackbar" class="ml-4 text-blue-300 hover:text-white focus:outline-none focus:text-white">
        &#10005;
      </button>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      show: Boolean,
      message: String,
      position: String,
      color: String
    },
    computed: {
      positionClasses() {
        return this.position === 'bottom' ? 'inset-x-0 bottom-0' : 'inset-x-0 top-0';
      },
      bgColor() {
        return `bg-${this.color}-500 || 'bg-gray-800'}`;
      }
    },
    methods: {
      closeSnackbar() {
        this.$emit("closeSnakeBar") ;
      }
    }
  };
  </script>