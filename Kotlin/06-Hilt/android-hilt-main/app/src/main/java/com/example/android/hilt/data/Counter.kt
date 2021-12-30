package com.example.android.hilt.data

import android.util.Log
import dagger.hilt.android.scopes.ActivityScoped
import javax.inject.Inject
import javax.inject.Singleton


class Counter @Inject constructor() : ICounter{
    private var value = 0

    override fun increase() {
        value += 1
    }

    override fun decrease() {
        value -= 1
    }

    override fun show() {
        Log.d("카운터값", value.toString())
    }

}