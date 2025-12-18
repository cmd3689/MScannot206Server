package framework

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/rs/zerolog/log"
)

type QuitHandler interface {
	Quit() error
}

type InputMachine struct {
	ctx        context.Context
	cancelFunc context.CancelFunc

	handler QuitHandler

	inputChan chan string
	doneChan  chan struct{}
	commands  map[string]ClientCommand
}

func (m *InputMachine) Init() {
	inputChan := make(chan string)
	m.inputChan = inputChan
	m.doneChan = make(chan struct{})

	m.commands = make(map[string]ClientCommand, 16)
}

func (m *InputMachine) AddCommand(cmd ClientCommand) error {
	if cmd == nil {
		return errors.New("cmd is nil")
	}

	for _, command := range cmd.Commands() {
		absC := strings.ToLower(command)
		if _, exists := m.commands[absC]; exists {
			log.Warn().Msgf("명령어 중복 등록 시도: %s", command)
			continue
		}
		m.commands[absC] = cmd
	}

	return nil
}

func (m *InputMachine) Attach(ctx context.Context, handler QuitHandler) {
	m.ctx, m.cancelFunc = context.WithCancel(ctx)
	m.handler = handler

	go m.taskCore()
}

func (m *InputMachine) Detach() {
	if m.cancelFunc != nil {
		m.cancelFunc()
	}
}

func (m *InputMachine) Done() <-chan struct{} {
	return m.ctx.Done()
}

func (m *InputMachine) taskCore() {
	go m.taskInput()

	for {
		select {
		case input := <-m.inputChan:
			m.handleInput(input)

		case <-m.ctx.Done():
			return
		}
	}
}

func (m *InputMachine) taskInput() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		tInput := strings.TrimSpace(input)
		if tInput == "" {
			continue
		}

		m.inputChan <- tInput
		<-m.doneChan
	}
}

func (m *InputMachine) handleInput(input string) {
	if strings.EqualFold(input, "-help") || strings.EqualFold(input, "-?") || strings.EqualFold(input, "-h") {
		m.printHelp()
		println()
		m.doneChan <- struct{}{}
	} else if strings.EqualFold(input, "-exit") || strings.EqualFold(input, "-quit") || strings.EqualFold(input, "-q") {
		if m.handler != nil {
			m.handler.Quit()
		}
	} else {
		if parts := strings.Fields(input); len(parts) > 0 {
			// parts[0]가 -로 시작하는지 확인
			if !strings.HasPrefix(parts[0], "-") {
				log.Error().Msgf("명령어는 '-'로 시작해야 합니다: %s", parts[0])
				fmt.Println("Usage: -help, -?, -h")
				println()
				m.doneChan <- struct{}{}
				return
			}

			cmdName := strings.ToLower(parts[0][1:])
			if cmd, exists := m.commands[cmdName]; exists {
				if err := cmd.Execute(parts[1:]); err != nil {
					log.Error().Msgf("명령어 실행 오류: %v", err)
					fmt.Println("Usage: -help, -?, -h")
					println()
				}
			} else {
				log.Error().Msgf("알 수 없는 명령어: %s", cmdName)
				fmt.Println("Usage: -help, -?, -h")
				println()
			}
		}
		m.doneChan <- struct{}{}
	}
}

func (m *InputMachine) printHelp() {
	fmt.Println("사용 가능한 명령어 목록:")
	println("-exit, -quit, -q: 프로그램 종료")

	keys := make([]string, 0, len(m.commands))
	for k := range m.commands {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for i := len(keys) - 1; i >= 0; i-- {
		key := keys[i]
		cmd := m.commands[key]

		fmt.Println()
		fmt.Println(cmd.Description())
	}
}
