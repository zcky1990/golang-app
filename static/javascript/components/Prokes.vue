<template>
    <section id="prokes">
        <div class="leading-4 container mx-auto">
            <div class="grid">
                <div class="grid p-4 gap-4">
                    <div class="title font-semibold text-xl text-center md:text-center lg:text-center xl:text-center">
                        Lorem Ipsum
                    </div>
                    <div
                        class="subtitle text-sm font-light text-justified md:text-center lg:text-center xl:text-center">
                        Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut
                        labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco
                        laboris nisi ut aliquip ex ea commodo consequat.
                    </div>
                </div>
                <div
                    class="grid gap-4 p-4 grid-cols-1 md:gap-4 lg:gap-4 xl:gap-4 md:grid-cols-2 lg:grid-cols-2 xl:grid-cols-2">
                    <div v-for="(item, index) in items" :key="index"
                        class="flex flex-col md:flex-row lg:flex-row xl:flex-row bg-white shadow-md">
                        <img class="h-48 w-full object-cover md:h-48 md:w-48 lg:h-48 lg:w-48 xl:h-48 xl:w-48"
                            :src="item.image" />
                        <div class="text-desc tracking-wide text-sm p-4">
                            <div class="uppercase tracking-wide text-sm font-semibold">{{ item.title }}</div>
                            <p class="mt-2 text-sm font-light tracking-wide text-slate-500">{{ item.description }}</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>

<script>
export default {
    name: 'Prokes',
    props: {
        items: {
            type: Array,
            default: () => [
                {
                    image: 'https://mdbootstrap.com/wp-content/uploads/2020/06/vertical.jpg',
                    title: 'Lorem Ipsum',
                    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.'
                },
                {
                    image: 'https://mdbootstrap.com/wp-content/uploads/2020/06/vertical.jpg',
                    title: 'Lorem Ipsum',
                    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.'
                },
                {
                    image: 'https://mdbootstrap.com/wp-content/uploads/2020/06/vertical.jpg',
                    title: 'Lorem Ipsum',
                    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.'
                },
                {
                    image: 'https://mdbootstrap.com/wp-content/uploads/2020/06/vertical.jpg',
                    title: 'Lorem Ipsum',
                    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.'
                }
            ]
        }
    }
}
</script>
