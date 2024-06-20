<template>
    <div :class="containerClass" aria-labelledby="modal-title" role="dialog" aria-modal="true">
        <div class="fixed inset-0 transition-opacity bg-gray-500 bg-opacity-75"></div>
        <div class="fixed inset-0 h-full mx-auto text-center">
            <div class="flex flex-col items-center justify-center w-full h-full max-w-lg mx-auto">
                <img decoding="async"
                    :src="imageUrl ? imageUrl : 'https://haribahagiaku.id/wp-content/uploads/2022/06/Tema-1-animasi-55-cosa.png'"
                    title="Tema-1-animasi-55-cosa" alt="Tema-1-animasi-55-cosa" />
                <h1 class="text-3xl font-bold text-white lg:text-4xl">{{ brideAndGroom }}</h1>
                <p class="mt-4 text-gray">Kepada Bapak/Ibu/Saudara/i</p>
                <h1 class="mt-4 text-3xl font-bold text-white lg:text-4xl">{{ invite }}</h1>
                <p class="mt-4 text-gray">Kami Mengundang Anda Untuk Hadir Di Acara Pernikahan Kami.</p>
                <div class="flex justify-center mt-2 space-x-2">
                    <div>
                        <button @click="handleClick"
                            class="relative inline-flex items-center px-8 py-3 overflow-hidden text-indigo-600 border border-current rounded group focus:outline-none focus:ring active:text-indigo-500"
                            href="#">
                            <span class="absolute transition-all -start-full group-hover:start-4">
                                <svg class="size-5 rtl:rotate-180" xmlns="http://www.w3.org/2000/svg" fill="none"
                                    viewBox="0 0 24 24" stroke="currentColor">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                        d="M17 8l4 4m0 0l-4 4m4-4H3" />
                                </svg>
                            </span>

                            <span class="text-sm font-medium transition-all group-hover:ms-4"> Download </span>
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
        imageUrl: {
            type: String,
            default: 'https://res.cloudinary.com/dk0z4ums3/image/upload/w_300,h_166,c_fill/v1718792056/attached_image/bentuk-perut-ibu-hamil-1-bulan-seperti-apa-ya.jpg'
        },
        name: {
            type: String,
            default: ''
        },
        invite: {
            type: String,
            default: ''
        },
        isEnvelopHidden: {
            type: Boolean,
            default: false
        },
        onClickButton: {
            type: Function,
            default: () => { }
        }
    },
    data() {
        return {
            brideAndGroom: 'Bride and Groom',
            isHidden: false
        };
    },
    computed: {
        containerClass() {
            return this.isHidden
                ? 'bg-white bg-cover bg-center h-screen hidden'
                : `bg-white bg-cover bg-center h-screen`;
        }
    },
    methods: {
        handleClick() {
            this.isHidden = true;
            this.onClickButton();
        }
    }
};
</script>