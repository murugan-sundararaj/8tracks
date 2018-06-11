# 8tracks mini version

## Services
### Tags service
    Schema:
        MAP of tag_name to tag_id
        MAP of tag_id to playlist_ids SET
        MAP of playlist_id to tag_ids SET
    
### Playlists service
    Schema:
         MAP of playlist_id to playlist object
         MAP of playlist_name to playlist_id

### Explore Service
    logic (.../explore/tagA+tagB)

    setA = fetch all playlist_ids from Tags Service that is tagged with tagA
    setB = fetch all playlist_ids from Tags Service that is tagged with tagB
    setC = intersection of setA and setB

    for each playlist_id in SetC
        SetD_1 thru SetD_n = fetch tag_ids from Tags service
     
    SetE = union of SetD_1 thru SetD_n // List<Tag>

    // List<Playlist>
    for each playlist_id in SetC
        fetch playlist object from Playlists service

    playlist is sorted by play count followed by likes count

## Flow
tag_id and playlist_id are UUIDs

**Create tagA**

    Request:
        curl --request POST \
        --url http://localhost:8080/tags/tag \
        --header 'content-type: application/json' \
        --data '{
        "tag_name": "tagA",
        "tag_type": 1
        }'

    Response:
        {"tag_id":"tag_uuid_1"}

**Create tagB**

    Request:
        curl --request POST \
        --url http://localhost:8080/tags/tag \
        --header 'content-type: application/json' \
        --data '{
        "tag_name": "tagB",
        "tag_type": 3
        }'

    Response:
        {"tag_id":"tag_uuid_2"}

**Create Playlist A**

specify the actual tag_ids created above 

    Request:
        curl --request POST \
        --url http://localhost:8080/playlists/playlist \
        --header 'content-type: application/json' \
        --data '{
        "playlist_name": "A",
        "tags": [
            {
                    "tag_id": "tag_uuid_1"
            },
            {
                    "tag_id": "tag_uuid_2"
            }
        ],
        "tracks": [
            {
                    "id": "trk_1",
                    "name": "ada ada"
            },
            {
                    "id": "trk_2",
                    "name": "blah blah"
            }
        ],
        "creator": {
            "id": "user_1",
            "name": "Tiger"
        }
        }'


    Response:
        {"playlist_id":"playlist_uuid_1"}

**Create Playlist B**

    Request:
        curl --request POST \
        --url http://localhost:8080/playlists/playlist \
        --header 'content-type: application/json' \
        --data '{
        "playlist_name": "B",
        "tags": [
            {
                    "tag_id": "tag_uuid_1"
            },
            {
                    "tag_id": "tag_uuid_2"
            }
        ],
        "creator": {
        "id": "user_1",
        "name": "Tiger"
        }
        }'

    Response:
        {"playlist_id":"playlist_uuid_2"}

**Explore**

    curl --request GET \
    --url 'http://localhost:8080/explore/tagA+tagB'

### APIs for CRUD on playlists, tags and explore service is available as Insomnia data dump 
        File: 8tracks.json
        
#### Insomnia - REST client
        https://insomnia.rest/download/

#### APIs
        Tags
        Load tag types
            GET {{ base_url  }}/tags/types
        Create a tag
            POST {{ base_url  }}/tags/tag
        Load a tag 
        by name
            GET {{ base_url  }}/tags/tag?names={tag_name_1;tag_name_2}
        by id
            GET {{ base_url  }}/tags/tag?ids={tag_id_1;tag_id_2}
        Remove a tag
            DELETE {{ base_url  }}/tags/{tag_id}
        Upsert a tag
            PUT {{ base_url  }}/tags/{tag_id}
        Update a tag
            PATCH {{ base_url  }}/tags/{tag_id}
        Tag playlist
            PUT {{ base_url  }}/tags/{tag_id}/playlists/{playlist_id}
        Untag playlist
            DELETE {{ base_url  }}/tags/{tag_id}/playlists/{playlist_id}

        Playlists
        Create a playlist
            POST {{ base_url  }}/playlists/playlist
        Load a playlist
        by name
            GET {{ base_url  }}/playlists/playlist?names={playlist_name_1;playlist_name_2}
        by id
            GET {{ base_url  }}/playlists/playlist?ids={playlist_id_1;playlist_id_2}
        Remove a playlist
            DELETE {{ base_url  }}/playlists/{playlist_id}
        Upsert a playlist
            PUT {{ base_url  }}/playlists/{playlist_id}
        Update playlist name
            PATCH {{ base_url  }}/playlists/{playlist_id}
            To update playlist's tag, refer Tags API (Tag playlist, UnTag playlist)
        Add track to a playlist
            PUT {{ base_url  }}/playlists/{playlist_id}/tracks/{track_id}
        Remove track off a playlist
            DELETE {{ base_url  }}/playlists/{playlist_id}/tracks/{track_id}
        Increment plays count
            POST {{ base_url  }}/playlists/{playlist_id}/plays
        Increment likes count
            POST {{ base_url  }}/playlists/{playlist_id}/likes
        Decrement likes count (dislike)
            POST {{ base_url  }}/playlists/{playlist_id}/dislikes

        Explore
            GET {{ base_url  }}/explore/tagA+tagB

### Build
```sh
    gb build
```
[gb](https://getgb.io/docs/install/) is used for building

### Unit Test    
```sh
    gb test -v
```

### Running the program
```sh
    ./bin/8tracks
        by default the program runs on port 8080
  
    ./bin/8tracks -listen.address :8585
        program runs on port 8585
```

### The program is written in Go
- Short variable names are common in Go

**Please let me know if any of the below things need to be implemented, I would be happy to do it.**

- All data is stored in memory. Did not want to create additional dependency by introducing datastore for the test problem
- It is not thread safe
- Pagination is not implemented
- Not validating the playlist id within the tag service
- Authentication and Authorization is not implemented
- No transaction. So, no rollback. Partial data is possible.
