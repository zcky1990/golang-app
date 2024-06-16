<template>
    <div class="relative flex min-h-screen ">
        <div
            class="flex flex-col items-center flex-auto min-w-0 bg-white sm:flex-row md:items-start sm:justify-center md:justify-start">
            <div class="relative items-center justify-start flex-auto hidden h-full p-10 overflow-hidden text-white bg-purple-900 bg-no-repeat bg-cover sm:w-1/2 xl:w-2/5 md:flex"
                style="background-image: url(https://images.unsplash.com/photo-1579451861283-a2239070aaa9?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&amp;ixlib=rb-1.2.1&amp;auto=format&amp;fit=crop&amp;w=1950&amp;q=80);">
                <div class="absolute inset-0 z-0 opacity-75 bg-gradient-to-b from-blue-900 to-gray-900"></div>
                <div class="absolute right-0 w-16 min-h-screen triangle" style=""></div>
                <img src="https://jasper-pimstorage-skullcandy.s3.us-west-1.amazonaws.com/bd2253a9671dac36a95faf821b52e78935050140be1718ce001f6aace45cf25c.png"
                    class="absolute mr-5 h-96 right-5">
                <div class="z-10 w-full max-w-md">
                    <div class="mb-6 font-bold leading-tight sm:text-4xl xl:text-5xl">Reference site about Lorem Ipsum..
                    </div>
                    <div class="font-normal text-gray-200 sm:text-sm xl:text-md"> What is Lorem Ipsum Lorem Ipsum is
                        simply dummy
                        text of the printing and typesetting industry Lorem Ipsum has been the industry's standard dummy
                        text ever
                        since the 1500s when an unknown printer took a galley of type and scrambled it to make a type
                        specimen book it
                        has?</div>
                </div>
                <ul class="circles">
                    <li></li>
                    <li></li>
                    <li></li>
                    <li></li>
                    <li></li>
                    <li></li>
                    <li></li>
                    <li></li>
                    <li></li>
                    <li></li>
                </ul>
            </div>
            <div
                class="w-2/5 w-full p-8 bg-white md:flex md:items-center md:justify-center sm:w-auto md:h-full xl:w-2/5 md:p-10 lg:p-14 sm:rounded-lg md:rounded-none ">
                <div class="w-full max-w-md space-y-8">
                    <div class="text-center">
                        <h2 class="mt-6 text-3xl font-bold text-gray-900">
                            Welcom Back!
                        </h2>
                        <p class="mt-2 text-sm text-gray-500">Please sign in to your account</p>
                    </div>
                    
                    <form class="mt-8 space-y-6" action="#" method="POST">
                        <input type="hidden" name="remember" value="true">
                        <div class="relative">
                            <div class="absolute mt-4 right-3"><svg xmlns="http://www.w3.org/2000/svg"
                                    class="w-6 h-6 text-green-500" fill="none" viewBox="0 0 24 24"
                                    stroke="currentColor">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                        d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                                </svg>
                            </div>
                            <label class="ml-3 text-sm font-bold tracking-wide text-gray-700">Email</label>
                            <input
                                class="w-full px-4 py-2 text-base border-b border-gray-300 focus:outline-none rounded-2xl focus:border-indigo-500"
                                type="" placeholder="mail@gmail.com">
                        </div>
                        <div class="content-center mt-8">
                            <label class="ml-3 text-sm font-bold tracking-wide text-gray-700">
                                Password
                            </label>
                            <input
                                class="content-center w-full px-4 py-2 text-base border-b border-gray-300 rounded-2xl focus:outline-none focus:border-indigo-500"
                                type="" placeholder="Enter your password">
                        </div>
                        <div class="flex items-center justify-between">
                            <div class="flex items-center">
                                <input id="remember_me" name="remember_me" type="checkbox"
                                    class="w-4 h-4 bg-blue-500 border-gray-300 rounded focus:ring-blue-400">
                                <label for="remember_me" class="block ml-2 text-sm text-gray-900">
                                    Remember me
                                </label>
                            </div>
                            <div class="text-sm">
                                <a href="#" class="text-indigo-400 hover:text-blue-500">
                                    Forgot your password?
                                </a>
                            </div>
                        </div>
                        <div>
                            <button type="submit"
                                class="flex justify-center w-full p-4 font-semibold tracking-wide text-gray-100 transition duration-500 ease-in rounded-full shadow-lg cursor-pointer bg-gradient-to-r from-indigo-500 to-blue-600 hover:bg-gradient-to-l hover:from-blue-500 hover:to-indigo-600">
                                Sign in
                            </button>
                        </div>
                        <p class="flex flex-col items-center justify-center mt-10 text-center text-gray-500 text-md">
                            <span>Don't have an account?</span>
                            <a href="#"
                                class="text-indigo-400 no-underline transition duration-300 ease-in cursor-pointer hover:text-blue-500 hover:underline">Sign
                                up</a>
                        </p>
                    </form>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
import { inject } from 'vue';

export default {
    created() {
        this.$store = inject('store');
    },
    methods: {

    }
}
</script>
