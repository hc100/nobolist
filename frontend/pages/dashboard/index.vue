<template>
  <div>
    <Nav />
    <header class="bg-white shadow">
      <div class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold text-gray-900">ダッシュボード</h1>
      </div>
    </header>
    <main class="pl-8">
      <div class="w-full max-w-xs"></div>
    </main>
  </div>
</template>

<script>
export default {
  apollo: {},
  middleware: ['isAuth'],
  data: () => ({}),
  head() {
    return {
      title: 'ダッシュボード',
    }
  },
  watch: {},
  mounted() {},
}
</script>
