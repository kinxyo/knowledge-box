# Nuxt as a frontend for Tauri

![NuxtTauriBlog](https://github.com/kinxyo/knowledge-box/assets/90744941/d79f3776-646e-4a71-ab16-b94e62321493)

## Introduction

**Nuxt** is a microframework for Vue.js. It is basically a fullstack framework for javascript. It includes both frontend & the backend, however we will only use the frontend part.

_One may be inclined to think that Nuxt might be an overkill, but it isn’t. Nuxt is much more than Vue. It’s provides a better frontend experience which compares to none. Vue, however, can be ideal for smaller application._

**Tauri**, on the other hand, is a toolkit for creating desktop application in Rust. The concept is similiar to Electron however Tauri requires lesser storage and memory (making it superior in some people’s book 👀).

## DEVELOPMENT

Here's how the folder structure is supposed to look:

```bash
.
|-- src
|   |-- assets
|   |-- components
|   |-- public
|   |-- layouts
|   |-- composables
|   |-- middleware
|   |-- pages
|   `-- app.vue
|-- src-tauri
|   |-- src
|   |-- build.rs
|   |-- Cargo.lock
|   |-- Cargo.toml
|   `-- tauri.conf.json
|-- tsconfig.json
|-- nuxt.config.ts
|-- package.json
|-- yarn.lock
`-- README.md
```

[Video Link](https://www.youtube.com/watch?v=MOnf_kGI6L0)

## PRODUCTION

For creating a binary for production, you must configure a few files; otherwise you may run into the below error.

![Top wtf moments when developing Tauri apps with Nuxt](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/muwd784uef5exx5hj9p6.png)

### Explanation

Usually, popular frontend frameworks like React and Vue, create a dedicated folder (i.e dist folder) where they keep the static build of the frontend by pre-rendering the code. These files are then served by the server and rendered by the browser. Since Tauri uses browser rendering engine to manifest the frontend therefore it naturally requires the static build. Hence, the above error occurs only when Tauri isn’t successful in finding the static build.

Nuxt, being a full-stack framework, is configured to by default render via its backend node server. This creates a conflict as Tauri requires static build files in advance to manifest the frontend of your application.

### Solution

Since we’re already using Rust for the backend therefore we’ll disable the server-side rendering for Nuxt.
Add the following code in _nuxt.config.ts_ :

```typescript
export default defineNuxtConfig({
    ssr: false
})
```

Now, we’ll simply tell Nuxt to prerender our site (meaning, create static files instead of relying on the server to run and process them during runtime). We do this by adding the following code to our _nuxt.config.ts_:

```typescript
routeRules: {
    '/': { prerender: true }
    /* do the same for all routes used */
},
```

Lastly, we need to specify the path to our static files folder. In tauri.conf.json, point to the public folder located in the .output folder.

```json
"build": {
    "distDir": "../.output/public"    
  },
```

### Completed

Cool! So now you can run the build command.

```bash
yarn run tauri build
```

If you aren’t using yarn then you can refer to other build commands [here](https://tauri.app/v1/guides/building/windows/).

## Closing Remark

To refer an actual project for better context, you can check out my github repo— [AI Therapist](https://github.com/kinxyo/CooperAI) application.
