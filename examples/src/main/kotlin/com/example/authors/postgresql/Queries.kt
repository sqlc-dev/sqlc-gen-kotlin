// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package com.example.authors.postgresql

import java.sql.Connection
import java.sql.SQLException
import java.sql.Statement

interface Queries {
  @Throws(SQLException::class)
  fun createAuthor(name: String, bio: String?): Author?
  
  @Throws(SQLException::class)
  fun deleteAuthor(id: Long)
  
  @Throws(SQLException::class)
  fun getAuthor(id: Long): Author?
  
  @Throws(SQLException::class)
  fun listAuthors(): List<Author>
  
}

