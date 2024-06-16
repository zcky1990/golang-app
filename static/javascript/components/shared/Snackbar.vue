<template>
  <div class="fixed w-full">
    <div v-if="show" class="max-w-md p-6 mx-auto">
      <div role="alert" :class="['p-4', bgColor, textColor, 'border', 'border-gray-100', 'rounded-xl']">
        <div class="flex items-start gap-4">
          <span v-if="!isError" class="text-green-600">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
              stroke="currentColor" class="w-6 h-6">
              <path stroke-linecap="round" stroke-linejoin="round"
                d="M9 12.75L11.25 15 15 9.75M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </span>
          <span v-else class="text-red-800">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-5 h-5">
              <path fill-rule="evenodd"
                d="M9.401 3.003c1.155-2 4.043-2 5.197 0l7.355 12.748c1.154 2-.29 4.5-2.599 4.5H4.645c-2.309 0-3.752-2.5-2.598-4.5L9.4 3.003zM12 8.25a.75.75 0 01.75.75v3.75a.75.75 0 01-1.5 0V9a.75.75 0 01.75-.75zm0 8.25a.75.75 0 100-1.5.75.75 0 000 1.5z"
                clip-rule="evenodd" />
            </svg>
          </span>

          <div class="flex-1">
            <strong :class="['block', 'font-medium', textColor]"> {{ title }} </strong>
            <strong class="block font-medium"> {{ message }} </strong>
          </div>

          <button @click="closeSnackbar" class="text-gray-500 transition hover:text-gray-600">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
              stroke="currentColor" class="w-6 h-6">
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
      </div>

    </div>
  </div>
</template>

<script>
export default {
  props: {
    show: Boolean,
    title: String,
    message: String,
    type: String,
    color: String,
    timeout: Number,
  },
  watch: {
    show: function (newVal, oldVal) { // watch it
      let self = this;
      if (newVal === true) {
        setTimeout(function () {
          self.closeSnackbar();
        }, self.timeout !== undefined ? self.timeout : 1000);
      }
    }
  },
  computed: {
    bgColor() {
      return this.type == 'info' ? 'bg-white' : 'bg-red-50';
    },
    textColor() {
      return this.type == 'info' ? 'text-gray-700' : 'text-red-700';
    },
    isError() {
      return this.type == 'error' ? true : false;
    }
  },
  methods: {
    closeSnackbar() {
      this.$emit("closeSnakeBar");
    }
  }
};
</script>