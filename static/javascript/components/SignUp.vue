<template>
    <div class="flex items-center justify-center ">
        <div class="container max-w-screen-lg mx-auto">
            <div class="p-4 px-4 mb-6 md:p-8">
                <div class="grid grid-cols-1 gap-4 text-sm gap-y-2 lg:grid-cols-3">
                    <h2 class="mt-10 text-2xl font-bold leading-9 tracking-tight text-gray-900">
                        Sign Up
                    </h2>

                    <div class="mt-10 lg:col-span-2">
                        <div class="grid grid-cols-1 gap-4 text-sm gap-y-2 md:grid-cols-5">
                            <div class="md:col-span-5">
                                <label for="full_name">Full Name</label>
                                <input type="text" name="full_name" id="full_name"
                                    class="w-full h-10 px-4 mt-1 border rounded bg-gray-50" value="" />
                            </div>

                            <div class="md:col-span-5">
                                <label for="email">Email Address</label>
                                <input type="text" name="email" id="email"
                                    class="w-full h-10 px-4 mt-1 border rounded bg-gray-50" value=""
                                    placeholder="email@domain.com" />
                            </div>

                            <div class="md:col-span-3">
                                <label for="address">Address / Street</label>
                                <input type="text" name="address" id="address"
                                    class="w-full h-10 px-4 mt-1 border rounded bg-gray-50" value="" placeholder="" />
                            </div>

                            <div class="md:col-span-2">
                                <label for="city">City</label>
                                <input type="text" name="city" id="city"
                                    class="w-full h-10 px-4 mt-1 border rounded bg-gray-50" value="" placeholder="" />
                            </div>

                            <div class="md:col-span-2">
                                <label for="country">Country / region</label>
                                <div class="flex items-center h-10 mt-1 border border-gray-200 rounded bg-gray-50">
                                    <input name="country" id="country" placeholder="Country"
                                        class="w-full px-4 text-gray-800 bg-transparent outline-none appearance-none"
                                        value="" />
                                    <button tabindex="-1"
                                        class="text-gray-300 transition-all outline-none cursor-pointer focus:outline-none hover:text-red-600">
                                        <svg class="w-4 h-4 mx-2 fill-current" xmlns="http://www.w3.org/2000/svg"
                                            viewBox="0 0 24 24" stroke="currentColor" stroke-width="2"
                                            stroke-linecap="round" stroke-linejoin="round">
                                            <line x1="18" y1="6" x2="6" y2="18"></line>
                                            <line x1="6" y1="6" x2="18" y2="18"></line>
                                        </svg>
                                    </button>
                                    <button tabindex="-1" for="show_more"
                                        class="text-gray-300 transition-all border-l border-gray-200 outline-none cursor-pointer focus:outline-none hover:text-blue-600">
                                        <svg class="w-4 h-4 mx-2 fill-current" xmlns="http://www.w3.org/2000/svg"
                                            viewBox="0 0 24 24" stroke="currentColor" stroke-width="2"
                                            stroke-linecap="round" stroke-linejoin="round">
                                            <polyline points="18 15 12 9 6 15"></polyline>
                                        </svg>
                                    </button>
                                </div>
                            </div>

                            <div class="md:col-span-2">
                                <label for="state">State / province</label>
                                <div class="flex items-center h-10 mt-1 border border-gray-200 rounded bg-gray-50">
                                    <input name="state" id="state" placeholder="State"
                                        class="w-full px-4 text-gray-800 bg-transparent outline-none appearance-none"
                                        value="" />
                                    <button tabindex="-1"
                                        class="text-gray-300 transition-all outline-none cursor-pointer focus:outline-none hover:text-red-600">
                                        <svg class="w-4 h-4 mx-2 fill-current" xmlns="http://www.w3.org/2000/svg"
                                            viewBox="0 0 24 24" stroke="currentColor" stroke-width="2"
                                            stroke-linecap="round" stroke-linejoin="round">
                                            <line x1="18" y1="6" x2="6" y2="18"></line>
                                            <line x1="6" y1="6" x2="18" y2="18"></line>
                                        </svg>
                                    </button>
                                    <button tabindex="-1" for="show_more"
                                        class="text-gray-300 transition-all border-l border-gray-200 outline-none cursor-pointer focus:outline-none hover:text-blue-600">
                                        <svg class="w-4 h-4 mx-2 fill-current" xmlns="http://www.w3.org/2000/svg"
                                            viewBox="0 0 24 24" stroke="currentColor" stroke-width="2"
                                            stroke-linecap="round" stroke-linejoin="round">
                                            <polyline points="18 15 12 9 6 15"></polyline>
                                        </svg>
                                    </button>
                                </div>
                            </div>

                            <div class="md:col-span-1">
                                <label for="zipcode">Zipcode</label>
                                <input type="text" name="zipcode" id="zipcode"
                                    class="flex items-center w-full h-10 px-4 mt-1 transition-all border rounded bg-gray-50"
                                    placeholder="" value="" />
                            </div>

                            <div class="md:col-span-5">
                                <div class="inline-flex items-center">
                                    <input type="checkbox" name="billing_same" id="billing_same"
                                        class="form-checkbox" />
                                    <label for="billing_same" class="ml-2">My billing address is different than
                                        above.</label>
                                </div>
                            </div>



                            <div class="text-right md:col-span-5">
                                <div class="inline-flex items-end">
                                    <button
                                        class="px-4 py-2 font-bold text-white bg-blue-500 rounded hover:bg-blue-700">Submit</button>
                                </div>
                            </div>

                        </div>
                    </div>
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