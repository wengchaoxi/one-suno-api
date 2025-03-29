# One Suno API

一个使用 Go 语言构建的 Suno API 网关服务。整合多个第三方 Suno API 提供商，通过加权轮询进行负载均衡，提供统一请求/响应格式的 Suno API 服务

**优势**

- *统一接口* - Suno AI 目前未开放官方 API，市面上存在多个基于爬虫实现的非官方 API 服务，One Suno API 提供标准化的请求/响应格式
- *负载均衡* - 使用加权轮询算法分配请求
- *易于扩展* - 模块化设计，支持快速接入新的第三方的 Suno API


## 接口文档

### `POST /v1/audio`

- Request

```json
{
    "is_custom": true, // 是否自定义，默认为 false
    "prompt": "[Verse]... [Verse 2]... [Chorus]...", // 歌词，当 is_custom 为 false 时，此字段为音乐描述
    "make_instrumental": false,                      // 是否纯音乐，默认为 false
    // 自定义模式下额外的参数
    "tags": "romantic ballad",  // 音乐类型
    "title": "a-romantic-song", // 音乐名称
    "continue_at": 180,         // 从音频 id 为 xxx 的音乐的 180 秒处继续生成，默认为 null
    "continue_clip_id": "xxx"   // 音频 id, 默认为 null
}
```

- Response

```json
{
    "code": 0,
    "msg": "Success",
    "data": [{
        "id": "6e81bebe-33ff-4527-a5c8-xxxxxxxxxxxx",
        "video_url": "https://cdn1.suno.ai/6e81bebe-33ff-4527-a5c8-xxxxxxxxxxxx.mp4",
        "audio_url": "https://cdn1.suno.ai/6e81bebe-33ff-4527-a5c8-xxxxxxxxxxxx.mp3",
        "image_url": "https://cdn1.suno.ai/image_6e81bebe-33ff-4527-a5c8-xxxxxxxxxxxx.png",
        "image_large_url": "https://cdn1.suno.ai/image_large_6e81bebe-33ff-4527-a5c8-xxxxxxxxxxxx.png",
        "major_model_version": "v3",
        "model_name": "chirp-v3",
        "metadata": {
            "tags": "romantic ballad",
            "prompt": "[Verse]... [Verse 2]... [Chorus]...",
            "gpt_description_prompt": null,
            "duration": 29.6
        },
        "created_at": "2024-04-09T08:35:53.414Z",
        "title": "a-romantic-song",
        "is_custom": true,
    }]
}
```


## 提供商列表

- [x] 接入 [acedata.cloud](https://platform.acedata.cloud/documents/suno-audios-integration/)
- [ ] 接入 [suno4.cn](https://suno4.cn/)
- [ ] 接入 [sunoapi.org](https://docs.sunoapi.org/)

---

项目基于 MIT 许可开源，欢迎 PR
