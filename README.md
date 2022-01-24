# green
 <!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol> 
    <li><a href="#run">Built with</a></li>
    <li><a href="#run">How to run app locally</a></li>  
     <li><a href="#run">API doc</a></li>  
      <li><a href="#run">Demo</a></li>  
  </ol>
</details>

## Display app:

 App => [Click Here to view app](https://green-code-assignment.herokuapp.com/).

## Built With:

Discover some packages of this project:

* [chi](https://pkg.go.dev/github.com/go-chi/chi/v5)
* [testing](https://pkg.go.dev/testing)
* [fmt](https://pkg.go.dev/fmt)
* [log](https://pkg.go.dev/log)
* [os](https://pkg.go.dev/os)
* [bufio](https://pkg.go.dev/bufio)
* [log](https://pkg.go.dev/log)
* [path](https://pkg.go.dev/path)
* [strings](https://pkg.go.dev/strings)
* [strconv](https://pkg.go.dev/strconv)
* [unicode](https://pkg.go.dev/unicode)

## Local setup:

### `#1`
For debug or run operation:
##### `git clone https://github.com/TaylanYildirim/green.git`
##### `cd green/ && go get && go run main.go`

1.
> Insert

```http
POST https://green-code-assignment.herokuapp.com/maze

```
| Parameter | Type | Description |
| :--- | :--- | :--- |
| `id` | `string` | **Required**. maze id |

2.
> Delete
```http

DELETE https://green-code-assignment.herokuapp.com/maze/{id}
e.g. https://green-code-assignment.herokuapp.com/maze/0

```



3.
> Get
```http

GET https://green-code-assignment.herokuapp.com/maze/{id}
e.g. https://green-code-assignment.herokuapp.com/maze/0

```

4.
> Update
```http

PUT https://green-code-assignment.herokuapp.com/maze/{id}
e.g. https://green-code-assignment.herokuapp.com/maze/0

```

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `id` | `string` | **Required**. Maze id |


## Demo:
### Insert Demo


https://user-images.githubusercontent.com/18633675/150860733-c9e4c088-c07b-482a-8e31-f00a3d3c5e0b.mov


### Delete Demo


https://user-images.githubusercontent.com/18633675/150860769-9a413326-0486-4834-ad6e-025b329a988d.mov


### Get Demo


https://user-images.githubusercontent.com/18633675/150860815-432e2906-4e66-4429-a8b8-75f58bfb2f99.mov


### Update Demo


https://user-images.githubusercontent.com/18633675/150860897-7aa15b9b-6b38-4301-9084-5296f780dde8.mov

