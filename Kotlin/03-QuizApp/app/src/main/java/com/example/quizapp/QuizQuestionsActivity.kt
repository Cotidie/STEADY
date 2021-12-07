package com.example.quizapp

import android.graphics.Color
import android.graphics.Typeface
import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.util.Log
import android.view.View
import android.widget.TextView
import androidx.core.content.ContextCompat
import com.example.quizapp.databinding.ActivityQuizQuestionsBinding

class QuizQuestionsActivity : AppCompatActivity(), View.OnClickListener {
    lateinit var binding: ActivityQuizQuestionsBinding
    private val questionList = Constants.getQeustions(5)
    private val tvOptions = ArrayList<TextView>()
    private var currentIndex: Int = 0
    private var selectedIndex: Int = 0


    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        binding = ActivityQuizQuestionsBinding.inflate(layoutInflater)

        initializeUI()
        setQuestion(0)
        setContentView(binding.root)
    }

    private fun initializeUI() {
        binding.progressBar.max = questionList.size
        tvOptions.add(binding.tvOptionOne)
        tvOptions.add(binding.tvOptionTwo)
        tvOptions.add(binding.tvOptionThree)
        tvOptions.add(binding.tvOptionFour)

        for (i in 0 until tvOptions.size) {
            tvOptions.get(i).setOnClickListener {
                selectedOptionView(i)
            }
        }
    }

    // idx번 퀴즈를 설정한다,
    private fun setQuestion(idx: Int) {
        val question = questionList.get(idx)
        binding.ivFlag.setImageResource(question.image)
        binding.progressBar.progress = idx+1
        binding.tvProgress.setText("${idx+1}/${binding.progressBar.max}")
        binding.tvQuestion.setText(question.question)
        binding.tvOptionOne.setText(question.optionOne)
        binding.tvOptionTwo.setText(question.optionTwo)
        binding.tvOptionThree.setText(question.optionThree)
        binding.tvOptionFour.setText(question.optionFour)

        if (idx == questionList.size-1) {
            binding.btnSubmit.text = "FINISH"
        }
    }

    // 옵션 TextView를 기본 스타일로 설정
    private fun defaultOptionsView() {
        for (option in tvOptions) {
            option.setTextColor(
                ContextCompat.getColor(
                    this, R.color.grey_dark
                )
            )
            option.background = ContextCompat.getDrawable(
                this, R.drawable.default_option_border_bg
            )
            option.typeface = Typeface.DEFAULT
        }
    }

    private fun selectedOptionView(selected: Int) {
        val selectedOption = tvOptions.get(selected)
        // 기본 버튼 스타일 초기화
        defaultOptionsView()
        // 선택 옵션 스타일 바꾸기
        selectedOption.setTextColor(Color.BLACK)
        selectedOption.background = ContextCompat.getDrawable(
            this, R.drawable.selected_option_border_bg
        )
        selectedOption.setTypeface(
            selectedOption.typeface,
            Typeface.BOLD
        )

        selectedIndex = selected
    }

    override fun onClick(v: View?) {
        Log.d("Activity Clicked: ", "Clicked!")
    }
}