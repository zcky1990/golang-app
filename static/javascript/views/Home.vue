<template>
    <contents />
      <router-view />
</template>
<script>
import contents from "./../components/HelloWorld.vue"

export default {
    name: 'App',
    components: {
        contents,
    }
}
</script>