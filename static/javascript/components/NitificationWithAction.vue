<template>
    <div v-if="show" class="modal fixed inset-0 z-10  overflow-auto bg-black bg-opacity-40">
        <div class="modal-content m-24 mx-auto w-4/5 grid h-3/4  place-content-center">
            <!-- <span class="close text-gray-500 float-right text-2xl font-bold cursor-pointer hover:text-black">&times;</span> -->
            <div class="max-w-md p-6 mx-auto h-fit">
                <div class="rounded-2xl border border-blue-100 bg-white p-4 shadow-lg sm:p-6 lg:p-8" role="alert">
                    <div class="flex items-center gap-4">
                        <span v-if="isNotification" class="shrink-0 rounded-full bg-blue-400 p-2 text-white">
                            <svg class="h-4 w-4" fill="currentColor" viewbox="0 0 20 20"
                                xmlns="http://www.w3.org/2000/svg">
                                <path clip-rule="evenodd"
                                    d="M18 3a1 1 0 00-1.447-.894L8.763 6H5a3 3 0 000 6h.28l1.771 5.316A1 1 0 008 18h1a1 1 0 001-1v-4.382l6.553 3.276A1 1 0 0018 15V3z"
                                    fill-rule="evenodd" />
                            </svg>
                        </span>

                        <span v-if="!isNotification" class="shrink-0 rounded-full bg-red-400 p-2 text-white">
                            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor"
                                class="w-5 h-5">
                                <path fill-rule="evenodd"
                                    d="M9.401 3.003c1.155-2 4.043-2 5.197 0l7.355 12.748c1.154 2-.29 4.5-2.599 4.5H4.645c-2.309 0-3.752-2.5-2.598-4.5L9.4 3.003zM12 8.25a.75.75 0 01.75.75v3.75a.75.75 0 01-1.5 0V9a.75.75 0 01.75-.75zm0 8.25a.75.75 0 100-1.5.75.75 0 000 1.5z"
                                    clip-rule="evenodd" />
                            </svg>
                        </span>

                        <p class="font-medium sm:text-lg">{{ titleDialog }}</p>
                    </div>

                    <p class="mt-4 text-gray-500 text-sm font-light">
                        {{ messageDialog }}
                    </p>

                    <div class="mt-6 sm:flex sm:gap-4 justify-center">
                        <button @click="handleSubmit"
                            :class="['inline-block w-2/4 rounded-lg bg-blue-500 px-5 py-3 text-center text-sm font-light sm:w-2/4', bgSubmitButton, textColorSubmitButton]"
                            href="#">
                            {{ button.submit }}
                        </button>

                        <button @click="handleCancel"
                            class="mt-2 inline-block w-2/4 rounded-lg bg-gray-50 px-5 py-3 text-center text-sm font-light text-gray-500 sm:mt-0 sm:w-2/4"
                            href="#">
                            {{ button.cancel }}
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </div>

</template>
<script>

export default {
    props: {
        show: {
            type: Boolean,
            default: true
        },
        titleDialog: {
            type: String,
            default: "Notification"
        },
        messageDialog: {
            type: String,
            default: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ipsam ea quo unde vel adipisci blanditiis voluptates eum. Nam, cum minima?"
        },
        button: {
            type: Object,
            default: () => ({
                submit: "Ok",
                cancel: "Cancel"
            })
        },
        typeDialog: {
            type: String,
            default: "warning"
        }
    },
    computed: {
        isNotification() {
            return this.typeDialog === 'notification' ? true : false;
        },
        bgSubmitButton() {
            if (this.typeDialog === 'notification') {
                return 'bg-blue-500'
            } else {
                return 'bg-red-400'
            }
        },
        textColorSubmitButton() {
            if (this.typeDialog === 'notification') {
                return 'text-white'
            } else {
                return 'text-white'
            }
        }
    },
    methods: {
        handleSubmit() {
            this.$emit("submitCallback");
        },
        handleCancel() {
            this.$emit("cancelCallback");
        }
    }
}
</script>