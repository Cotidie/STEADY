package com.example.roomwithview.room.repository

import androidx.annotation.WorkerThread
import com.example.roomwithview.room.dao.WordDao
import com.example.roomwithview.room.entity.Word
import kotlinx.coroutines.flow.Flow

// Declares the DAO as a private property in the constructor. Pass in the DAO
// instead of the whole database, because you only need access to the DAO
class WordRepository(private val wordDao: WordDao) {

    // Room executes all queries on a separate thread.
    // Observed Flow will notify the observer when the data has changed.
    // Observer는 대체 어떻게 등록되는 것인지
    val allWords: Flow<List<Word>> = wordDao.getAlphabetizedWords()

    
    // By default Room runs suspend queries off the main thread, therefore, we don't need to
    // implement anything else to ensure we're not doing long running database work
    // off the main thread.
    @Suppress("RedundantSuspendModifier")
    @WorkerThread
    suspend fun insert(word: Word) {
        wordDao.insert(word)
    }
}