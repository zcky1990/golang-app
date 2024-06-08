<template>
    <div class="container mx-auto">
        <h1 class="text-3xl font-bold underline">{{ this.$store.state.count }}</h1>
        <h1 class="text-3xl font-bold underline">{{ this.$store.state.data }}</h1>
    </div>
</template>
<script>
import { inject } from 'vue';

export default {
    props: ['dataProperty'],
    data() {
        return {
            message: 'Hello, Vue! JS AJA 12358910',
        }
    },
    created() {
        this.$store = inject('store');
    }
}
</script>