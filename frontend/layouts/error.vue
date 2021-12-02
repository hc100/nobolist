<template>
  <div>
    <Header />
    <main>
      <div>
        <div class="w-full max-w-xs">
          <div class="bg-green-200 rounded px-6 pt-6 pb-6 mb-4">
            {{ pageTitle }}
          </div>
        </div>
        <div class="w-full max-w-xs">
          <div class="bg-green-200 rounded px-6 pt-6 pb-6 mb-4">
            <p class="read" v-html="errorMessage"></p>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
export default {
  props: {
    error: {
      type: Object,
      default: null,
    },
  },
  data() {
    return {
      pageTitle: 'Error',
    }
  },
  head() {
    return {
      title: this.pageTitle,
    }
  },
  computed: {
    errorMessage() {
      if (this.error.statusCode === 404) {
        return '申し訳ありません。<br>お探しのページが見つかりませんでした。'
      }
      if (this.error.message) {
        return this.error.message
      }
      return 'エラーが発生しました'
    },
  },
  beforeMount() {},
  created() {
    if (this.error.statusCode === 404) {
      this.pageTitle = '404 File Not Found'
    } else if (this.error.statusCode === 500) {
      this.pageTitle = '500 Internal Server Error'
    } else {
      this.pageTitle = 'Server Error'
    }
  },
}
</script>
