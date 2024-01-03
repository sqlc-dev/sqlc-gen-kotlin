// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package com.example.booktest.postgresql

import java.time.OffsetDateTime

enum class BookType(val value: String) {
  FICTION("FICTION"),
  NONFICTION("NONFICTION");

  companion object {
    private val map = BookType.values().associateBy(BookType::value)
    fun lookup(value: String) = map[value]
  }
}

data class Author (
  val authorId: Int,
  val name: String
)

data class Book (
  val bookId: Int,
  val authorId: Int,
  val isbn: String,
  val bookType: BookType,
  val title: String,
  val year: Int,
  val available: OffsetDateTime,
  val tags: List<String>
)

