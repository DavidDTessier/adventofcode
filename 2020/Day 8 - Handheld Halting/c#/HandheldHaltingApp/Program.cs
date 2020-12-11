using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Reflection;

namespace HandheldHaltingApp
{
    class Program
    {
        const string AccumulatorStep = "acc";
        const string JumpStep = "jmp";
        const string NoOpStep = "nop";


        static void Main(string[] args)
        {
            Console.WriteLine("Hello World!");
            var ins = LoadInstructions();
       
            var stepsRan = new List<int>();

            
            var flow = AccumulatorCount(0, 0, ins.ToList(), stepsRan);

            Console.Write($"Accumulator Count Before Any Step Is Executed A Second Time: {flow.Accumulator}");
            stepsRan.Clear();

            flow = SwapActionsAndAccumulate(0, 0, ins.ToList(), stepsRan);

            Console.Write($"Accumulator Count After the program terminates {flow.Accumulator}");
        }

        static InstructionFlow AccumulatorCount(int stepIndex,int accumCount, List<CodeInstruction> actions, List<int> stepsRan)
        {

            while (stepIndex != actions.Count)
            {
                var step = actions[stepIndex];

                if (stepsRan.Contains(stepIndex))
                    return new InstructionFlow { Accumulator = accumCount, InfiniteLoopHit = true };

                stepsRan.Add(stepIndex);

                if (step.Action == AccumulatorStep)
                {
                    accumCount += step.StepCount;
                    stepIndex++;
                }
                else if (step.Action == JumpStep)
                {
                    stepIndex += step.StepCount;
                }
                else if (step.Action == NoOpStep)
                {
                    stepIndex++;
                }


            }

            return new InstructionFlow { Accumulator = accumCount, InfiniteLoopHit = false };
        }

        static InstructionFlow SwapActionsAndAccumulate(int stepIndex, int accumCount, List<CodeInstruction> actions, List<int> stepsRan)
        {

            while (stepIndex != actions.Count)
            {
                var step = actions[stepIndex];

                if (step.Action == AccumulatorStep)
                {
                    stepsRan.Add(stepIndex);
                    accumCount += step.StepCount;
                    stepIndex++;
                }
                else
                {
                    var swapStep = new CodeInstruction { Action = step.Action == JumpStep ? NoOpStep : JumpStep, StepCount = step.StepCount };
                    actions[stepIndex] = swapStep;

  
                    var flow = AccumulatorCount(stepIndex, accumCount, actions, stepsRan.ConvertAll(idx => idx).ToList());

                    if (flow.InfiniteLoopHit)
                    {
                        actions[stepIndex] = step;
                        stepsRan.Add(stepIndex);
                        stepIndex += step.Action == JumpStep ? step.StepCount : 1;
                    }
                    else
                    {
                        flow.FlowFixMessage = $"Replace Action {step.Action} with {swapStep.Action} at idx: {stepIndex} fixes operation flow.";
                        return flow;
                    }
                }


            }

            return new InstructionFlow { Accumulator = accumCount, InfiniteLoopHit = false }; 
        }

      

        static IList<CodeInstruction> LoadInstructions()
        {
            var inputFilePath = Path.GetDirectoryName(Assembly.GetCallingAssembly().Location)
                + Path.DirectorySeparatorChar + "input.dat";

            var instructions = File.ReadAllLines(inputFilePath).ToList().Select(l => 
                new CodeInstruction { Action = l.Split(" ")[0].Trim(), StepCount = int.Parse(l.Split(" ")[1].Trim()) }
            ).ToList();
            return instructions;
        }

    }

    class CodeInstruction
    {
        public string Action { get; set; }
        public int StepCount { get; set; }
    }

    class InstructionFlow
    {
        public bool InfiniteLoopHit { get; set; }
        public int Accumulator { get; set; }
        public string FlowFixMessage { get; set; }
    }
}
