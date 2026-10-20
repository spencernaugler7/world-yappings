# World Sayings Archive

### setup

install the [just](https://github.com/casey/just#installation) cmdline utility

then run this command in root dir

```bash
just install-deps
```

### Dependancy Docs

- [gin](https://gin-gonic.com/en/docs/)
- [gorm](https://gorm.io/docs/)
- [htmx](https://four.htmx.org/docs)
- [blades css](https://blades.dev/#docs)

### Tools

- [live reload](https://github.com/air-verse/air)

## Planning

Individual record is a `Saying`

with these attributes

- `Original`
    - saying in it's original language
- `Translation`
    - literal translation to english
- `Meaning`
    - english meaning of the saying
- `When to use` (Optional)
- `Example` (Optional)


There are `Tags` to categorize all the `Saying`s

with these attributes

- `Name`
    - name of the tag 


Cartesian record to link tags to sayings.


### Functionality

- [ ] create  a saying 
- [ ] edit a saying
    - [ ] ability to add tags 
- [ ] ability to search for sayings 
    - [ ] search by title by default
    - [ ] Be able to search by tag