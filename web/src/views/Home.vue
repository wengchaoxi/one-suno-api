<template>
<div>
    <h1>管理后台</h1>
    <nav>
    <router-link to="/providers">Providers 管理</router-link> |
    <router-link to="/api-keys">API Keys 管理</router-link>
    </nav>
    <router-view />
</div>
</template>
