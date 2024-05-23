<template>
    <div class="container mx-auto">
        <h1 class="text-3xl font-bold underline">{{ message }}</h1>
        <h1 class="text-3xl underline">{{ dataProperty.JsImportPath }}</h1>
        <h1 class="text-3xl underline">{{ dataProperty.CssImportPath }}</h1>
    </div>
</template>
<script>
export default {
    props: ['dataProperty'],
    data() {
        return {
            message: 'Hello, Vue! JS AJA 12358910',
        }
    }
}
</script>