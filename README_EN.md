# One Suno API

**English** | [简体中文](./README.md)

## Overview

A Suno API gateway service built with Go. Integrates multiple third-party Suno API providers, provides load balancing through weighted round-robin, and offers unified request/response format for Suno API services.

**Key Features**

- Unified Interface: With no official API from Suno AI, One Suno API standardizes the fragmented ecosystem of third-party APIs by providing a consistent interface.
- Load Balancing: Uses weighted round-robin algorithm to distribute requests.
- Easy to Extend: Modular design supports quick integration of new third-party Suno APIs.
- Containerized Deployment: Easy deployment and management using Docker.

## API Documentation

### `POST /v1/audio`

- Request

```json
{
    "is_custom": true, // Whether to use custom mode, defaults to false
    "prompt": "[Verse]... [Verse 2]... [Chorus]...", // Lyrics when is_custom is true, otherwise music description
    "make_instrumental": false,                      // Whether instrumental only, defaults to false
    // Additional parameters in custom mode
    "tags": "romantic ballad",  // Music tags
    "negative_tags": "violence", // Negative tags, defaults to null
    "title": "a-romantic-song", // Music title
    "continue_at": 180,         // Continue generation from 180s of the audio with id xxx, defaults to null
    "continue_clip_id": "xxx"   // Audio id, defaults to null
}
```

- Response

```json
{
    "code": 0,
    "msg": "Success",
    "data": [{
        "id": "6e81bebe-33ff-4527-a5c8-xxxxxxxxxxxx",
        "title": "a-romantic-song",
        "video_url": "https://cdn1.suno.ai/6e81bebe-33ff-4527-a5c8-xxxxxxxxxxxx.mp4",
        "audio_url": "https://cdn1.suno.ai/6e81bebe-33ff-4527-a5c8-xxxxxxxxxxxx.mp3",
        "image_url": "https://cdn1.suno.ai/image_6e81bebe-33ff-4527-a5c8-xxxxxxxxxxxx.png",
        "image_large_url": "https://cdn1.suno.ai/image_large_6e81bebe-33ff-4527-a5c8-xxxxxxxxxxxx.png",
        "major_model_version": "v3",
        "model_name": "chirp-v3",
        "metadata": {
            "tags": "romantic ballad",
            "prompt": "a romantic song",
            "gpt_description_prompt": "a romantic song",
            "lyric": "[Verse]... [Verse 2]... [Chorus]...",
            "duration": 29.6
        },
        "created_at": "2024-04-09T08:35:53.414Z"
    }]
}
```

## Supported Providers

- [x] [acedata.cloud](https://platform.acedata.cloud/documents/suno-audios-integration/)
- [ ] [piapi.ai](https://piapi.ai/docs/music-api/create-task)
- [ ] [sunoapi.org](https://docs.sunoapi.org/)

---

This project is open source under the MIT License. PRs are welcome!
